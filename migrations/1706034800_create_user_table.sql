-- +goose Up
CREATE TYPE UserType AS ENUM ('USER', 'ADMIN');

-- Create "user" table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    telegram_id BIGINT UNIQUE,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR,
    username VARCHAR,
    language TEXT DEFAULT 'en',
    user_type UserType NOT NULL DEFAULT 'USER',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- Create triggers for updating "updated_at" column in various tables
CREATE TRIGGER update_user_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- +goose Down
DROP TABLE IF EXISTS "User";
DROP TYPE IF EXISTS UserType;