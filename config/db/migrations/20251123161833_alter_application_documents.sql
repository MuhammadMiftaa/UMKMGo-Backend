-- +goose Up
-- +goose StatementBegin
ALTER TABLE application_documents ALTER COLUMN type TYPE text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE application_documents ALTER COLUMN type TYPE document_type;
-- +goose StatementEnd
