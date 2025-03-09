BEGIN;

CREATE TABLE users
(
    id         UUID PRIMARY KEY    NOT NULL,
    name       VARCHAR(100)        NOT NULL CHECK (LENGTH(name) >= 3),
    email      VARCHAR(255) UNIQUE NOT NULL,
    password   TEXT                NOT NULL CHECK (LENGTH(password) >= 8),
    role_id    UUID                NOT NULL,
    created_at TIMESTAMP           NOT NULL,
    updated_at TIMESTAMP           NOT NULL,
    deleted_at TIMESTAMP           NULL
);

CREATE UNIQUE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_deleted_at ON users (deleted_at);

COMMIT;
