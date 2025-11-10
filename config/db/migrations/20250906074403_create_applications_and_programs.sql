-- +goose Up
-- +goose StatementBegin
CREATE TYPE program_type AS ENUM ('training', 'certification', 'funding');
CREATE TYPE training_type AS ENUM ('online', 'offline', 'hybrid');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE programs (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  description TEXT,
  banner TEXT,
  provider VARCHAR(100),
  provider_logo TEXT,
  type program_type NOT NULL, -- training / certification / funding
  training_type training_type, -- hanya untuk training & certification
  batch INT,
  batch_start_date DATE,
  batch_end_date DATE,
  location VARCHAR(100),
  min_amount NUMERIC(15,2),
  max_amount NUMERIC(15,2),
  interest_rate NUMERIC(5,2),
  max_tenure_months INT,
  application_deadline DATE,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_by INT REFERENCES users(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ,
  CONSTRAINT programs_training_fields CHECK (
    (type IN ('training','certification') AND (
        (training_type IS NOT NULL) OR (training_type IS NULL)
    )) OR (type = 'funding')
  )
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE program_benefits (
  id SERIAL PRIMARY KEY,
  program_id INT NOT NULL REFERENCES programs(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE program_requirements (
  id SERIAL PRIMARY KEY,
  program_id INT NOT NULL REFERENCES programs(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS program_requirements;
DROP TABLE IF EXISTS program_benefits;
DROP TABLE IF EXISTS programs;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TYPE IF EXISTS training_type;
DROP TYPE IF EXISTS program_type;
-- +goose StatementEnd
