CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE unit_kind AS ENUM (
    'NHQ',
    'REGION',
    'WING',
    'GROUP',
    'SQUADRON',
    'FLIGHT'
);

CREATE TYPE unit_category AS ENUM (
    'ADMIN',
    'COMPOSITE',
    'CADET',
    'SENIOR'
);

CREATE TABLE units (
    charter_number  VARCHAR(11) PRIMARY KEY,
    kind            unit_kind,
    category        unit_category,
    name            TEXT NOT NULL,
    address         TEXT NOT NULL,
    city            TEXT NOT NULL
);