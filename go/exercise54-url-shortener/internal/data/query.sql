-- name: GetOneURL :one
SELECT * FROM urls
WHERE short_url = ?
LIMIT 1;

-- name: AllURLs :many
SELECT * FROM urls
ORDER BY id;

-- name: CreateURL :one
INSERT INTO urls (short_url, long_url)
VALUES (?, ?)
RETURNING *;

-- name: UpdateURL :exec
UPDATE urls
SET short_url = ?, long_url = ?
WHERE id = ?
RETURNING *;

-- name: DeleteURL :exec
DELETE FROM urls
WHERE id = ?;
