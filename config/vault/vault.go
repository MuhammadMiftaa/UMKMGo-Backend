package vault

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"UMKMGo-backend/config/env"
	"UMKMGo-backend/config/log"

	vault "github.com/hashicorp/vault/api"
)

var VaultClient *vault.Client

// ~ NewVaultClientWithAppRole membuat Vault client dan login AppRole.
// ~ vaultAddr contoh: "https://vault.example.com:8200"
func SetupVault(vaultCfg env.Vault) {
	cfg := vault.DefaultConfig()
	cfg.Address = vaultCfg.Addr
	// Jika perlu custom TLS config, set cfg.HttpClient.Transport etc.

	client, err := vault.NewClient(cfg)
	if err != nil {
		log.Log.Fatalf("failed create vault client: %v", err)
	}

	// Login AppRole
	data := map[string]any{
		"role_id":   vaultCfg.RoleID,
		"secret_id": vaultCfg.SecretID,
	}
	// auth/approle/login
	secret, err := client.Logical().Write("auth/approle/login", data)
	if err != nil {
		log.Log.Fatalf("approle login failed: %v", err)
	}
	if secret == nil || secret.Auth == nil || secret.Auth.ClientToken == "" {
		log.Log.Fatalf("no token returned from approle login")
	}

	client.SetToken(secret.Auth.ClientToken)
	// optional: you may want to set token accessor for renew or background renew
	VaultClient = client
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
