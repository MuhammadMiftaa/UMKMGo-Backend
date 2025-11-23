-- +goose Up
-- +goose StatementBegin
ALTER TABLE application_histories ALTER COLUMN actioned_by DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE application_histories ALTER COLUMN actioned_by SET NOT NULL;
-- +goose StatementEnd
