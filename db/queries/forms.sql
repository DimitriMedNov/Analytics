-- name: CreateQuestion :one
INSERT INTO questions (
    question, answer_type
) VALUES ($1, $2) RETURNING *;

-- name: CreateQuestionWithOpts :one
INSERT INTO questions (
    question, answer_type, options
) VALUES ($1, $2, $3) RETURNING *;
