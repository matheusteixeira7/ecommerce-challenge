BEGIN;

CREATE TABLE roles
(
    id         UUID PRIMARY KEY   NOT NULL,
    name       VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP          NOT NULL,
    updated_at TIMESTAMP          NOT NULL
);

COMMIT;
