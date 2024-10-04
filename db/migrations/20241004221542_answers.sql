-- migrate:up
CREATE TABLE IF NOT EXISTS answers
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    value TEXT NOT NULL,
    form_id BIGINT NOT NULL,
    question_id BIGINT NOT NULL,
    FOREIGN KEY (form_id) REFERENCES forms (
        id
    ) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (question_id) REFERENCES questions (
        id
    ) ON UPDATE CASCADE ON DELETE CASCADE

);

-- migrate:down
DROP TABLE IF EXISTS answers;
