
-- +migrate Up

CREATE TABLE IF NOT EXISTS teams (
    id VARCHAR(50),
    "name" VARCHAR(100) NOT NULL,
    "type" VARCHAR(100) NOT NULL UNIQUE,
    hub_id  VARCHAR(50) NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,

    CONSTRAINT teams_pkey PRIMARY KEY (id),
    CONSTRAINT fk_team_hub FOREIGN KEY (hub_id) REFERENCES hubs(id)
);

-- +migrate Down

DROP TABLE IF EXISTS teams;