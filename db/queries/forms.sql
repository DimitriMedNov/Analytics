-- name: CreateQuestion :one
INSERT INTO questions (question, answer_type, options)
VALUES ($1, $2, $3)
RETURNING *;


-- name: CreateAnswer :one
INSERT INTO answers (id, value, form_id, question_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateAnswer :one
UPDATE answers
SET value = $2, form_id = $3, question_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAnswer :one
DELETE FROM answers
WHERE id = $1
RETURNING *;

-- name: CreateForm :one
INSERT INTO forms (id, name, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateForm :one
UPDATE forms
SET name = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteForm :one
DELETE FROM forms
WHERE id = $1
RETURNING *;

-- name: CreateFormQuestion :one
INSERT INTO forms_questions (form_id, question_id, required, depends_on)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateFormQuestion :one
UPDATE forms_questions
SET required = $3, depends_on = $4
WHERE form_id = $1 AND question_id = $2
RETURNING *;

-- name: DeleteFormQuestion :one
DELETE FROM forms_questions
WHERE form_id = $1 AND question_id = $2
RETURNING *;
