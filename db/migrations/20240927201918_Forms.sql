-- migrate:up
CREATE TABLE IF NOT EXISTS forms
(
    id          BIGSERIAL    NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    CONSTRAINT forms_pkey PRIMARY KEY (id)

);
-- migrate:down

DROP TABLE IF EXISTS forms;
