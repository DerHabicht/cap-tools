CREATE TYPE user_role AS ENUM (
    'ROOT',
    'COMMAND',
    'ADMIN',
    'MEMBER',
    'PARENT'
);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT NOT NULL UNIQUE
);

CREATE TABLE user_unit_roles (
    unit_charter VARCHAR(11) NOT NULL REFERENCES units (charter_number) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    role user_role NOT NULL,
    PRIMARY KEY(unit_charter, user_id)
);

CREATE TABLE user_member_links (
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    member_capid INT NOT NULL REFERENCES members (capid) ON DELETE CASCADE,
    self BOOLEAN NOT NULL DEFAULT TRUE,
    PRIMARY KEY(user_id, member_capid)
);