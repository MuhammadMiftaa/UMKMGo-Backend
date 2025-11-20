-- +goose Up
-- +goose StatementBegin
ALTER TABLE umkms 
ADD COLUMN revenue_record TEXT,
ADD COLUMN business_permit TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE umkms 
DROP COLUMN IF EXISTS revenue_record,
DROP COLUMN IF EXISTS business_permit;
-- +goose StatementEnd