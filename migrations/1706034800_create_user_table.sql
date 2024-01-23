-- +goose Up
CREATE TYPE UserType AS ENUM ('USER', 'ADMIN');

CREATE TABLE IF NOT EXISTS "User" (
    id         TEXT PRIMARY KEY DEFAULT gen_random_uuid(),
    telegramID BIGINT UNIQUE,
    firstName  TEXT NOT NULL,
    lastName   TEXT NOT NULL,
    username   TEXT DEFAULT 'none',
    language   TEXT DEFAULT 'en',
    userType   UserType NOT NULL DEFAULT 'USER',
    createdAt  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- +goose Down
DROP TABLE IF EXISTS "User";
DROP TYPE IF EXISTS UserType;