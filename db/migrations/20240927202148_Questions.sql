-- migrate:up
CREATE TABLE IF NOT EXISTS questions (
    id BIGSERIAL NOT NULL,
    question TEXT NOT NULL,
    answer_type VARCHAR(255) NOT NULL,
    options TEXT[],
    CONSTRAINT questions_pkey PRIMARY KEY (id)

);

-- migrate:down
DROP TABLE IF EXISTS questions;
