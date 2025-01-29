-- +goose Up

-- Get a random string
-- source : https://www.depesz.com/2017/02/06/generate-short-random-textual-ids/
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_random_string(
        IN string_length INTEGER,
        IN possible_chars TEXT
        DEFAULT '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
    ) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
    output TEXT = '';
    i INT4;
    pos INT4;
BEGIN
    FOR i IN 1..string_length LOOP
        pos := 1 + CAST( random() * ( LENGTH(possible_chars) - 1) AS INT4 );
        output := output || substr(possible_chars, pos, 1);
    END LOOP;
    RETURN output;
END;
$$;
-- +goose StatementEnd


CREATE TYPE UserType AS ENUM ('USER', 'ADMIN');

-- Create "user" table
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY DEFAULT get_random_string(8),
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
