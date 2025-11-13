-- +goose Up
-- +goose StatementBegin
CREATE TABLE slas (
    id SERIAL PRIMARY KEY,
    status VARCHAR(50) NOT NULL UNIQUE,
    max_days INT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT check_status CHECK (status IN ('screening', 'final'))
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO slas (status, max_days, description) VALUES
('screening', 7, 'Maksimal 7 hari untuk proses screening'),
('final', 14, 'Maksimal 14 hari untuk keputusan akhir');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS slas;
-- +goose StatementEnd