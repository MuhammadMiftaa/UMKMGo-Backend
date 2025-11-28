-- +goose Up
-- +goose StatementBegin
ALTER TABLE umkms ALTER COLUMN nib TYPE TEXT;
ALTER TABLE umkms ALTER COLUMN npwp TYPE TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE umkms ALTER COLUMN nib TYPE VARCHAR(50);
ALTER TABLE umkms ALTER COLUMN npwp TYPE VARCHAR(50);
-- +goose StatementEnd
