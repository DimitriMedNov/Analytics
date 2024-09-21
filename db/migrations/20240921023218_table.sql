-- migrate:up
CREATE TABLE IF NOT EXISTS locations (
    id BIGSERIAL PRIMARY KEY,
    lat BIGINT NOT NULL,
    long BIGINT NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS locations;
