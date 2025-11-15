-- +goose Up
-- +goose StatementBegin
CREATE TYPE otp_status AS ENUM ('active', 'used');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE otps (
  id BIGSERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL, -- email of the user
  phone_number VARCHAR(255) NULL, -- phone number of the user
  otp_code VARCHAR(255) NOT NULL, -- one-time password
  temp_token TEXT NULL, -- temporary token
  status otp_status NOT NULL DEFAULT 'active',
  expires_at TIMESTAMP WITH TIME ZONE NOT NULL, -- expiration time for the OTP
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE NULL,
  deleted_at TIMESTAMP WITH TIME ZONE NULL,

  CONSTRAINT temp_token_unique UNIQUE (temp_token)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS otps;
DROP TYPE IF EXISTS otp_status;
-- +goose StatementEnd
