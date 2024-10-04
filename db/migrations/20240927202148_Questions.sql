-- migrate:up
CREATE TYPE answer_type AS ENUM ('text', 'number', 'multiple_choice', 'single_choice');
CREATE TABLE IF NOT EXISTS questions
(
    id          BIGSERIAL    NOT NULL PRIMARY KEY,
    question    TEXT         NOT NULL,
    answer_type answer_type NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS questions;
