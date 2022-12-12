
-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(50),
    "name" VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    team_id  VARCHAR(50) NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,

    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT fk_user_team FOREIGN KEY (team_id) REFERENCES teams(id)
);

-- +migrate Down

DROP TABLE IF EXISTS users;