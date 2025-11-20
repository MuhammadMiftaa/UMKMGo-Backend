-- +goose Up
-- +goose StatementBegin
CREATE TABLE training_applications (
    id SERIAL PRIMARY KEY,
    application_id INT NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    motivation TEXT NOT NULL,
    business_experience TEXT,
    learning_objectives TEXT,
    availability_notes TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE certification_applications (
    id SERIAL PRIMARY KEY,
    application_id INT NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    business_sector VARCHAR(100) NOT NULL, -- Sektor usaha: F&B, Fashion, Craft, dll
    product_or_service VARCHAR(255) NOT NULL, -- Produk/layanan yang akan disertifikasi
    business_description TEXT NOT NULL, -- Deskripsi usaha
    years_operating INT, -- Berapa tahun usaha berjalan
    current_standards TEXT, -- Standar/prosedur yang sudah diterapkan
    certification_goals TEXT NOT NULL, -- Tujuan mendapatkan sertifikasi
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE funding_applications (
    id SERIAL PRIMARY KEY,
    application_id INT NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    business_sector VARCHAR(100) NOT NULL,
    business_description TEXT NOT NULL,
    years_operating INT,
    requested_amount NUMERIC(15,2) NOT NULL,
    fund_purpose TEXT NOT NULL,
    business_plan TEXT,
    revenue_projection NUMERIC(15,2),
    monthly_revenue NUMERIC(15,2),
    requested_tenure_months INT NOT NULL,
    collateral_description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS training_applications;
DROP TABLE IF EXISTS certification_applications;
DROP TABLE IF EXISTS funding_applications;
-- +goose StatementEnd
