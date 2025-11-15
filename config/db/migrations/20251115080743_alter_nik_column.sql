-- +goose Up
-- +goose StatementBegin
ALTER TABLE umkms ALTER COLUMN nik SET DATA TYPE TEXT;
ALTER TABLE umkms ALTER COLUMN kartu_number SET DATA TYPE TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE umkms ALTER COLUMN nik SET DATA TYPE VARCHAR(255);
ALTER TABLE umkms ALTER COLUMN kartu_number SET DATA TYPE VARCHAR(255);
-- +goose StatementEnd
