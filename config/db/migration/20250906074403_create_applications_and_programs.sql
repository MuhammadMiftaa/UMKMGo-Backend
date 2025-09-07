-- +goose Up
-- NOTE: Perbaikan file ini: mengganti kurung kurawal ke kurung biasa, menambahkan tabel programs yang benar, mengganti kolom training_id -> program_id, dan menambahkan bagian Down lengkap.
-- +goose StatementBegin
CREATE TYPE program_type AS ENUM ('training', 'certification', 'funding');
CREATE TYPE application_status AS ENUM ('draft', 'submitted', 'screening', 'revision', 'assessment', 'decision', 'approved', 'rejected', 'cancelled');
CREATE TYPE application_stage AS ENUM ('screening', 'assessment', 'decision', 'execution', 'completed');
CREATE TYPE document_type AS ENUM ('ktp', 'nib', 'npwp', 'proposal', 'portfolio', 'rekening', 'other');
CREATE TYPE training_type AS ENUM ('online', 'offline', 'hybrid');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE applications (
  id SERIAL PRIMARY KEY,
  umkm_id INT NOT NULL REFERENCES umkms(id) ON DELETE CASCADE,
  type program_type NOT NULL,
  status application_status NOT NULL DEFAULT 'draft',
  current_stage application_stage,
  score NUMERIC(5,2),
  submitted_at TIMESTAMPTZ,
  decided_at TIMESTAMPTZ,
  notes TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION applications_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER set_applications_update_timestamp
BEFORE UPDATE ON applications
FOR EACH ROW
EXECUTE FUNCTION applications_update_timestamp();
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE application_documents (
  id SERIAL PRIMARY KEY,
  application_id INT NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
  type document_type NOT NULL,
  file_path VARCHAR(255) NOT NULL,
  file_name VARCHAR(255) NOT NULL,
  verified_by INT REFERENCES users(id),
  verified_at TIMESTAMPTZ,
  is_valid BOOLEAN DEFAULT FALSE,
  remarks TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION application_documents_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER set_application_documents_update_timestamp
BEFORE UPDATE ON application_documents
FOR EACH ROW
EXECUTE FUNCTION application_documents_update_timestamp();
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE application_reviews (
  id SERIAL PRIMARY KEY,
  application_id INT NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
  reviewer_id INT NOT NULL REFERENCES users(id),
  comments TEXT,
  stage application_stage NOT NULL,
  action VARCHAR(50) NOT NULL,
  reason_code VARCHAR(50),
  remarks TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION application_reviews_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER set_application_reviews_update_timestamp
BEFORE UPDATE ON application_reviews
FOR EACH ROW
EXECUTE FUNCTION application_reviews_update_timestamp();
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
CREATE OR REPLACE FUNCTION programs_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER set_programs_update_timestamp
BEFORE UPDATE ON programs
FOR EACH ROW
EXECUTE FUNCTION programs_update_timestamp();
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
CREATE OR REPLACE FUNCTION program_benefits_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER set_program_benefits_update_timestamp
BEFORE UPDATE ON program_benefits
FOR EACH ROW
EXECUTE FUNCTION program_benefits_update_timestamp();
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

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION program_requirements_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER set_program_requirements_update_timestamp
BEFORE UPDATE ON program_requirements
FOR EACH ROW
EXECUTE FUNCTION program_requirements_update_timestamp();
-- +goose StatementEnd

-- +goose Down
-- Urutan: drop trigger -> function -> table -> type (reverse dependency)

-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_program_requirements_update_timestamp ON program_requirements;
DROP TRIGGER IF EXISTS set_program_benefits_update_timestamp ON program_benefits;
DROP TRIGGER IF EXISTS set_programs_update_timestamp ON programs;
DROP TRIGGER IF EXISTS set_application_reviews_update_timestamp ON application_reviews;
DROP TRIGGER IF EXISTS set_application_documents_update_timestamp ON application_documents;
DROP TRIGGER IF EXISTS set_applications_update_timestamp ON applications;
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS program_requirements_update_timestamp();
DROP FUNCTION IF EXISTS program_benefits_update_timestamp();
DROP FUNCTION IF EXISTS programs_update_timestamp();
DROP FUNCTION IF EXISTS application_reviews_update_timestamp();
DROP FUNCTION IF EXISTS application_documents_update_timestamp();
DROP FUNCTION IF EXISTS applications_update_timestamp();
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS program_requirements;
DROP TABLE IF EXISTS program_benefits;
DROP TABLE IF EXISTS programs;
DROP TABLE IF EXISTS application_reviews;
DROP TABLE IF EXISTS application_documents;
DROP TABLE IF EXISTS applications;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TYPE IF EXISTS training_type;
DROP TYPE IF EXISTS document_type;
DROP TYPE IF EXISTS application_stage;
DROP TYPE IF EXISTS application_status;
DROP TYPE IF EXISTS program_type;
-- +goose StatementEnd
