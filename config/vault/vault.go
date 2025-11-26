package vault

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"UMKMGo-backend/config/env"
	"UMKMGo-backend/config/log"

	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/types/model"
	"UMKMGo-backend/internal/utils"

	vault "github.com/hashicorp/vault/api"
)

type DecryptParams struct {
	UserID    int
	UMKMID    *int
	FieldName string
	TableName string
	RecordID  int
	Purpose   string
	IPAddress string
	UserAgent string
	RequestID string
}

var VaultClient *vault.Client
var vaultConfig env.Vault

// ~ NewVaultClientWithAppRole membuat Vault client dan login AppRole.
func SetupVault(vaultCfg env.Vault) {
	vaultConfig = vaultCfg
	
	cfg := vault.DefaultConfig()
	cfg.Address = vaultCfg.Addr

	client, err := vault.NewClient(cfg)
	if err != nil {
		log.Log.Fatalf("failed create vault client: %v", err)
	}

	// Login AppRole
	if err := loginAppRole(client); err != nil {
		log.Log.Fatalf("approle login failed: %v", err)
	}

	VaultClient = client

	// Start auto renewal
	go autoRenewToken()
	
	log.Log.Info("Vault client initialized with auto-renewal")
}

// ~ loginAppRole melakukan login ke Vault menggunakan AppRole.
func loginAppRole(client *vault.Client) error {
	data := map[string]interface{}{
		"role_id":   vaultConfig.RoleID,
		"secret_id": vaultConfig.SecretID,
	}

	secret, err := client.Logical().Write("auth/approle/login", data)
	if err != nil {
		return err
	}
	if secret == nil || secret.Auth == nil || secret.Auth.ClientToken == "" {
		return errors.New("no token returned from approle login")
	}

	client.SetToken(secret.Auth.ClientToken)
	log.Log.Infof("Vault login successful, token TTL: %d seconds", secret.Auth.LeaseDuration)
	return nil
}

// ~ autoRenewToken secara otomatis memperbarui token Vault setiap 30 menit.
func autoRenewToken() {
	ticker := time.NewTicker(30 * time.Minute) // Renew setiap 30 menit
	defer ticker.Stop()

	for range ticker.C {
		// Coba renew token
		_, err := VaultClient.Auth().Token().RenewSelf(3600)
		if err != nil {
			log.Log.Warnf("Token renewal failed: %v, re-authenticating...", err)
			// Jika gagal, login ulang
			if err := loginAppRole(VaultClient); err != nil {
				log.Log.Errorf("Re-authentication failed: %v", err)
			} else {
				log.Log.Info("Re-authentication successful")
			}
		} else {
			log.Log.Debug("Token renewed successfully")
		}
	}
}

// ~ EncryptTransit encrypts plaintext using Transit key.
// ~ transitKey is the name of the transit key in Vault (e.g. "nik-key").
// ~ transitMount is the mount path of the transit engine, usually "transit".
// ~ Returns ciphertext string (e.g. "vault:v1:...") which is safe to store.
func EncryptTransit(ctx context.Context, transitMount, transitKey string, plaintext []byte) (string, error) {
	if VaultClient == nil {
		return "", errors.New("vault client is nil")
	}
	if transitMount == "" {
		transitMount = "transit"
	}

	// base64 encode plaintext, Transit expects base64 plaintext
	b64Plain := base64.StdEncoding.EncodeToString(plaintext)

	path := fmt.Sprintf("%s/encrypt/%s", transitMount, transitKey)
	req := map[string]any{
		"plaintext": b64Plain,
		// "context": "<base64 optional>", // optional: context for key derivation / additional authenticated data
		// "key_version": 2, // optional: force specific key version
	}

	secret, err := VaultClient.Logical().WriteWithContext(ctx, path, req)
	if err != nil {
		return "", fmt.Errorf("vault encrypt error: %w", err)
	}
	if secret == nil || secret.Data == nil {
		return "", errors.New("empty response from transit/encrypt")
	}
	ctIface, ok := secret.Data["ciphertext"]
	if !ok {
		return "", errors.New("ciphertext missing in response")
	}
	ciphertext, ok := ctIface.(string)
	if !ok {
		return "", errors.New("ciphertext is not string")
	}
	return ciphertext, nil
}

// ~ DecryptTransit decrypts ciphertext using Transit key.
// ~ Returns plaintext bytes.
func DecryptTransit(ctx context.Context, transitMount, transitKey, ciphertext string) ([]byte, error) {
	if VaultClient == nil {
		return nil, errors.New("vault client is nil")
	}
	if transitMount == "" {
		transitMount = "transit"
	}
	path := fmt.Sprintf("%s/decrypt/%s", transitMount, transitKey)
	req := map[string]interface{}{
		"ciphertext": ciphertext,
		// "context": "<base64 optional>", // must match encrypt if used
	}

	secret, err := VaultClient.Logical().WriteWithContext(ctx, path, req)
	if err != nil {
		return nil, fmt.Errorf("vault decrypt error: %w", err)
	}
	if secret == nil || secret.Data == nil {
		return nil, errors.New("empty response from transit/decrypt")
	}
	plainB64Iface, ok := secret.Data["plaintext"]
	if !ok {
		return nil, errors.New("plaintext missing in response")
	}
	plainB64, ok := plainB64Iface.(string)
	if !ok {
		return nil, errors.New("plaintext not string")
	}
	plaintext, err := base64.StdEncoding.DecodeString(plainB64)
	if err != nil {
		return nil, fmt.Errorf("failed decode plaintext base64: %w", err)
	}
	return plaintext, nil
}

// DecryptWithLog decrypts data and logs to vault_decrypt_logs
func DecryptWithLog(
	ctx context.Context,
	ciphertext string,
	encryptionKey string,
	params DecryptParams,
	vaultLogRepo repository.VaultDecryptLogRepository,
) (string, error) {
	// Perform decryption
	plaintext, err := DecryptTransit(
		ctx,
		env.Cfg.Vault.TransitPath,
		encryptionKey,
		ciphertext,
	)

	// Create log entry
	logEntry := model.VaultDecryptLog{
		UserID:    params.UserID,
		UMKMID:    params.UMKMID,
		FieldName: params.FieldName,
		TableName: params.TableName,
		RecordID:  params.RecordID,
		Purpose:   params.Purpose,
		IPAddress: params.IPAddress,
		UserAgent: params.UserAgent,
		RequestID: params.RequestID,
		Success:   err == nil,
	}

	if err != nil {
		logEntry.ErrorMessage = err.Error()
	}

	// Log the decrypt operation (don't fail if logging fails)
	if logErr := vaultLogRepo.LogDecrypt(ctx, logEntry); logErr != nil {
		log.Log.Errorf("failed to log decrypt operation: %v", logErr)
	}

	return utils.MaskMiddle(string(plaintext)), err
}

// DecryptNIKWithLog is a helper for NIK decryption
func DecryptNIKWithLog(
	ctx context.Context,
	ciphertext string,
	params DecryptParams,
	vaultLogRepo repository.VaultDecryptLogRepository,
) (string, error) {
	params.FieldName = "nik"
	params.TableName = "umkms"
	return DecryptWithLog(ctx, ciphertext, env.Cfg.Vault.NIKEncryptionKey, params, vaultLogRepo)
}

// DecryptKartuNumberWithLog is a helper for Kartu Number decryption
func DecryptKartuNumberWithLog(
	ctx context.Context,
	ciphertext string,
	params DecryptParams,
	vaultLogRepo repository.VaultDecryptLogRepository,
) (string, error) {
	params.FieldName = "kartu_number"
	params.TableName = "umkms"
	return DecryptWithLog(ctx, ciphertext, env.Cfg.Vault.KartuEncryptionKey, params, vaultLogRepo)
}

// GetContextInfo extracts common context information
func GetContextInfo(ctx context.Context) (ipAddress, userAgent, requestID string) {
	if val := ctx.Value("ipAddress"); val != nil {
		ipAddress, _ = val.(string)
	}
	if val := ctx.Value("userAgent"); val != nil {
		userAgent, _ = val.(string)
	}
	if val := ctx.Value("requestID"); val != nil {
		requestID, _ = val.(string)
	}
	return
}
