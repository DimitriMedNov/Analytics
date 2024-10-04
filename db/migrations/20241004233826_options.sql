-- migrate:up
CREATE TABLE IF NOT EXISTS options
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    value TEXT NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS options;
