-- name: GetURL :one
SELECT * FROM urls
WHERE id = ? LIMIT 1;

-- name: AllURLs :many
SELECT * FROM authons
ORDER BY name;

-- name: CreateURL :one
INSERT INTO urls (short_url, long_url) 
VALUES (?, ?)
RETURNING *;

-- name: UpdateURL :exec
UPDATE urls
set short_url = ?, long_url = ?
WHERE id = ?
RETURNING *;

-- name: DeleteURL :exec
DELETE FROM urls
WHERE id = ?;
