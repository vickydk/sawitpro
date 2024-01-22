BEGIN;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone_number    varchar(13) NOT NULL,
    name            varchar(64) NOT NULL,
    password        VARCHAR(255) NOT NULL,
    success_login   int  NOT NULL default 0,
    created_at      TIMESTAMP        DEFAULT NOW(),
    updated_at      TIMESTAMP        DEFAULT NOW()
);

COMMIT;