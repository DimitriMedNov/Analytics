-- name: GetAllLocationsByUser :many
SELECT l.id, l.lat, l.long FROM locations l
JOIN locations_users ON locations_users.location_id = locations.id 
WHERE users_id = $1;

-- name: CreateLocation :one
INSERT INTO locations (lat, long) VALUES ($1, $2) RETURNING *;

