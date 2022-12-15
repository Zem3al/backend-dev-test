
-- +migrate Up

CREATE TABLE IF NOT EXISTS hubs (
    id VARCHAR(50),
    "name" VARCHAR(100) NOT NULL,
    location VARCHAR(500) UNIQUE,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,

    CONSTRAINT hubs_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS hubs_name_index ON hubs (name);

-- +migrate Down

DROP TABLE IF EXISTS hubs;