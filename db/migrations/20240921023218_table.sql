-- migrate:up
CREATE TABLE IF NOT EXISTS locations
(
    id   BIGSERIAL PRIMARY KEY,
    lat  NUMERIC(14, 7) NOT NULL,
    long NUMERIC(14, 7) NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS locations;
