CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE unit_kind AS ENUM (
    'nhq',
    'region',
    'wing',
    'group',
    'squadron',
    'flight'
);

CREATE TYPE unit_category AS ENUM (
    'admin',
    'composite',
    'cadet',
    'senior'
);

CREATE TABLE units (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    deleted_at      TIMESTAMP,
    charter_number  VARCHAR(11) NOT NULL UNIQUE,
    kind            unit_kind,
    category        unit_category,
    name            TEXT NOT NULL,
    address         TEXT NOT NULL,
    city            TEXT NOT NULL
);

CREATE INDEX deleted_units ON units (deleted_at);