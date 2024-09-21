-- migrate:up
CREATE TABLE IF NOT EXISTS locations_users (
    location_id BIGINT NOT NULL,
    users_id BIGINT NOT NULL,
    CONSTRAINT fk_location FOREIGN KEY (location_id)
    REFERENCES locations (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
    CONSTRAINT fk_users FOREIGN KEY (users_id)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

-- migrate:down
DROP DATABASE IF EXISTS lcoations_users;
