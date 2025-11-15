# Step-by-Step Implementasi Vault Transit untuk Enkripsi NIK dengan Golang

## 1. **Setup Vault Transit Engine**

### A. Login ke Vault UI
```
URL: https://vault.miftech.web.id
```

### B. Enable Transit Secret Engine (via UI)
1. Login ke Vault UI
2. Klik **"Secrets"** di menu atas
3. Klik **"Enable new engine"**
4. Pilih **"Transit"**
5. Isi konfigurasi:
   - **Path**: `transit` (atau `nik-encryption`)
   - **Description**: "Transit engine for NIK encryption"
6. Klik **"Enable Engine"**

### C. Create Encryption Key (via UI)
1. Masuk ke Transit engine yang baru dibuat
2. Klik **"Create encryption key"**
3. Isi konfigurasi:
   - **Name**: `nik-key` (atau `personal-data-key`)
   - **Type**: `aes256-gcm96` (recommended untuk data sensitif)
   - **Derived**: ☑️ Check (untuk context-based encryption)
   - **Exportable**: ☐ Uncheck (lebih aman)
   - **Allow plaintext backup**: ☐ Uncheck
4. Klik **"Create encryption key"**

### D. Alternative: Setup via Vault CLI
```bash
# Enable transit engine
vault secrets enable transit

# Create encryption key
vault write -f transit/keys/nik-key \
  type=aes256-gcm96 \
  derived=true \
  exportable=false

# Verify key
vault read transit/keys/nik-key
```

## 2. **Setup Authentication & Policy**

### A. Create Policy for Application (via UI)
1. Klik **"Policies"** di menu atas
2. Klik **"Create ACL policy"**
3. **Name**: `nik-encryption-policy`
4. **Policy** (paste ini):

```hcl
# Policy untuk enkripsi/dekripsi NIK
path "transit/encrypt/nik-key" {
  capabilities = ["update"]
}

path "transit/decrypt/nik-key" {
  capabilities = ["update"]
}

path "transit/keys/nik-key" {
  capabilities = ["read"]
}

# Untuk key rotation (optional)
path "transit/keys/nik-key/rotate" {
  capabilities = ["update"]
}

# Untuk rewrap (optional - saat key rotation)
path "transit/rewrap/nik-key" {
  capabilities = ["update"]
}
```

5. Klik **"Create policy"**

### B. Create AppRole for Golang App (via UI)
1. Klik **"Access"** di menu atas
2. Klik **"Enable new method"**
3. Pilih **"AppRole"**
4. **Path**: `approle`
5. Klik **"Enable Method"**

6. Setelah enabled, klik pada **approle**
7. Klik **"Create role"**
8. Konfigurasi:
   - **Role name**: `nik-app-role`
   - **Policies**: Pilih `nik-encryption-policy`
   - **Secret ID TTL**: `720h` (30 hari)
   - **Token TTL**: `1h`
   - **Token Max TTL**: `24h`
   - **Bind Secret ID**: ☑️ Yes
9. Klik **"Save"**

# Setup AppRole via Vault CLI atau API

Karena Vault UI tidak mendukung konfigurasi lengkap AppRole, kita akan menggunakan **Vault CLI** atau **API** untuk setup.

## Opsi 1: Setup via Vault CLI (Recommended)

### A. Install Vault CLI

```bash
# Linux/WSL
wget https://releases.hashicorp.com/vault/1.15.4/vault_1.15.4_linux_amd64.zip
unzip vault_1.15.4_linux_amd64.zip
sudo mv vault /usr/local/bin/

# macOS
brew install vault

# Windows (PowerShell as Admin)
choco install vault

# Verify installation
vault --version
```

### B. Configure Vault CLI untuk Connect ke Server

```bash
# Set Vault address
export VAULT_ADDR='https://vault.miftech.web.id'
export VAULT_TOKEN=xxxxxxxxxxxxxxxxxxxxxxx

# Login ke Vault (gunakan token dari UI)
vault login
# Masukkan token yang didapat dari Vault UI
```

### F. Get Role ID

```bash
# Get Role ID
vault read auth/approle/role/nik-app-role/role-id

# Output:
# Key        Value
# ---        -----
# role_id    a1b2c3d4-e5f6-7890-abcd-ef1234567890
```

**Simpan Role ID ini!**

### G. Generate Secret ID

```bash
# Generate Secret ID
vault write -f auth/approle/role/nik-app-role/secret-id

# Output:
# Key                   Value
# ---                   -----
# secret_id             f1e2d3c4-b5a6-9807-cdef-1234567890ab
# secret_id_accessor    abcdef12-3456-7890-abcd-ef1234567890
# secret_id_ttl         720h
```

**Simpan Secret ID ini dengan AMAN! Secret ID hanya ditampilkan sekali.**

### H. Generate Multiple Secret IDs (Optional)

```bash
# Generate dengan metadata untuk tracking
vault write auth/approle/role/nik-app-role/secret-id \
    metadata="environment=production,app=nik-service"

# Generate dengan custom TTL
vault write auth/approle/role/nik-app-role/secret-id \
    ttl=168h

# List semua Secret ID Accessors (untuk revoke nanti)
vault list auth/approle/role/nik-app-role/secret-id
```

### I. Test Authentication

```bash
# Test login dengan AppRole
vault write auth/approle/login \
    role_id="YOUR_ROLE_ID" \
    secret_id="YOUR_SECRET_ID"

# Jika berhasil, akan return token:
# Key                     Value
# ---                     -----
# token                   hvs.CAESID...
# token_accessor          abc123...
# token_duration          1h
# token_renewable         true
# token_policies          ["default" "nik-encryption-policy"]
```

---

## Update .env File

Setelah mendapatkan **Role ID** dan **Secret ID**, update file `.env`
