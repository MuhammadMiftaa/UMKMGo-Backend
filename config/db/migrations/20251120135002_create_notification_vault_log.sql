-- Migration 1: Create notifications table
-- File: config/db/migrations/YYYYMMDDHHMMSS_create_notifications.sql

-- +goose Up
-- +goose StatementBegin
CREATE TYPE notification_type AS ENUM (
    'application_submitted',
    'screening_approved',
    'screening_rejected', 
    'screening_revised',
    'final_approved',
    'final_rejected',
    'program_reminder',
    'document_required',
    'general_info'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    umkm_id INT NOT NULL REFERENCES umkms(id) ON DELETE CASCADE,
    application_id INT REFERENCES applications(id) ON DELETE SET NULL,
    type notification_type NOT NULL,
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMPTZ,
    metadata JSONB, -- Additional data (program info, admin name, etc)
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_notifications_umkm_id ON notifications(umkm_id);
CREATE INDEX idx_notifications_is_read ON notifications(is_read);
CREATE INDEX idx_notifications_created_at ON notifications(created_at DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_notifications_created_at;
DROP INDEX IF EXISTS idx_notifications_is_read;
DROP INDEX IF EXISTS idx_notifications_umkm_id;
DROP TABLE IF EXISTS notifications;
DROP TYPE IF EXISTS notification_type;
-- +goose StatementEnd

---

-- Migration 2: Create vault decrypt log table
-- File: config/db/migrations/YYYYMMDDHHMMSS_create_vault_decrypt_log.sql

-- +goose Up
-- +goose StatementBegin
CREATE TYPE decrypt_purpose AS ENUM (
    'profile_view',
    'application_review',
    'application_creation',
    'profile_update',
    'admin_verification',
    'report_generation',
    'compliance_audit',
    'system_process'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE vault_decrypt_logs (
    id BIGSERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    umkm_id INT REFERENCES umkms(id) ON DELETE SET NULL,
    field_name VARCHAR(100) NOT NULL, -- 'nik', 'kartu_number', etc
    table_name VARCHAR(100) NOT NULL, -- 'umkms', etc
    record_id INT NOT NULL, -- ID of the decrypted record
    purpose decrypt_purpose NOT NULL,
    ip_address INET,
    user_agent TEXT,
    request_id VARCHAR(50), -- For tracking related operations
    success BOOLEAN NOT NULL DEFAULT TRUE,
    error_message TEXT,
    decrypted_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_vault_decrypt_logs_user_id ON vault_decrypt_logs(user_id);
CREATE INDEX idx_vault_decrypt_logs_umkm_id ON vault_decrypt_logs(umkm_id);
CREATE INDEX idx_vault_decrypt_logs_decrypted_at ON vault_decrypt_logs(decrypted_at DESC);
CREATE INDEX idx_vault_decrypt_logs_purpose ON vault_decrypt_logs(purpose);
CREATE INDEX idx_vault_decrypt_logs_field_name ON vault_decrypt_logs(field_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_vault_decrypt_logs_field_name;
DROP INDEX IF EXISTS idx_vault_decrypt_logs_purpose;
DROP INDEX IF EXISTS idx_vault_decrypt_logs_decrypted_at;
DROP INDEX IF EXISTS idx_vault_decrypt_logs_umkm_id;
DROP INDEX IF EXISTS idx_vault_decrypt_logs_user_id;
DROP TABLE IF EXISTS vault_decrypt_logs;
DROP TYPE IF EXISTS decrypt_purpose;
-- +goose StatementEnd