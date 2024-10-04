-- migrate:up
CREATE TABLE IF NOT EXISTS question_option (
    option_id BIGINT NOT NULL,
    question_id BIGINT NOT NULL,
    FOREIGN KEY (option_id) REFERENCES options (
        id
    ) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (question_id) REFERENCES questions (
        id
    ) ON UPDATE CASCADE ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS question_option;
