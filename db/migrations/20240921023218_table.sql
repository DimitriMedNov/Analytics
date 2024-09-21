-- migrate:up
CREATE TABLE IF NOT EXISTS locations (
    id BIGSERIAL PRIMARY KEY,
    lat NUMERIC(10, 8),
    long NUMERIC(10, 8)
);

-- migrate:down
DROP TABLE IF EXISTS locations;

