-- +goose Up
-- +goose StatementBegin
ALTER TABLE umkms ADD COLUMN photo text, ADD COLUMN qr_code text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE umkms DROP COLUMN photo, DROP COLUMN qr_code;
-- +goose StatementEnd
