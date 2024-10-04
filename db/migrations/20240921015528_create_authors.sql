-- migrate:up
CREATE TABLE IF NOT EXISTS users
(
    id    BIGSERIAL PRIMARY KEY,
    name  TEXT NOT NULL,
    admin BOOL DEFAULT false
);

-- migrate:down
DROP TABLE IF EXISTS users;
