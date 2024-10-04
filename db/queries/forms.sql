-- name: CreateQuestion :one
INSERT INTO questions (question, answer_type)
VALUES ($1, $2)
RETURNING *;


