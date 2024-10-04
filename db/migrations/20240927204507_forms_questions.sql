-- migrate:up
CREATE TABLE IF NOT EXISTS forms_questions
(
    form_id     BIGINT NOT NULL,
    question_id BIGINT NOT NULL,
    required    BOOL   NOT NULL DEFAULT FALSE,
    depends_on  BIGINT NULL,
    FOREIGN KEY (depends_on) references questions (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_form FOREIGN KEY (form_id)
        REFERENCES forms (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_question FOREIGN KEY (question_id)
        REFERENCES questions (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);


-- migrate:down

DROP TABLE IF EXISTS forms_questions;