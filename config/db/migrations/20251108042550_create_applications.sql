-- +goose Up
-- +goose StatementBegin
CREATE TYPE application_status AS ENUM (
    'screening',
    'revised',
    'final',
    'approved',
    'rejected'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TYPE application_history_action AS ENUM (
    'submit',
    'revise',
    'approve_by_admin_screening',
    'reject_by_admin_screening',
    'approve_by_admin_vendor',
    'reject_by_admin_vendor'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TYPE application_type AS ENUM ('training', 'certification', 'funding');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TYPE document_type AS ENUM ('ktp', 'nib', 'npwp', 'proposal', 'portfolio', 'rekening', 'other');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    umkm_id INT NOT NULL REFERENCES umkms(id) ON DELETE CASCADE,
    program_id INT NOT NULL REFERENCES programs(id) ON DELETE CASCADE,
    type application_type NOT NULL,
    status application_status NOT NULL DEFAULT 'screening',
    submitted_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMP
)
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE application_documents (
    id SERIAL PRIMARY KEY,
    application_id INT NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    type document_type NOT NULL,
    file TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMP
)
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE application_histories (
    id SERIAL PRIMARY KEY,
    application_id INT NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    status application_history_action NOT NULL,
    notes TEXT,
    actioned_at TIMESTAMPTZ DEFAULT NOW(),
    actioned_by INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE application_histories;
DROP TABLE application_documents;
DROP TABLE applications;
DROP TYPE IF EXISTS application_type;
DROP TYPE IF EXISTS document_type;
DROP TYPE IF EXISTS application_status;
DROP TYPE IF EXISTS application_history_action;
-- +goose StatementEnd
