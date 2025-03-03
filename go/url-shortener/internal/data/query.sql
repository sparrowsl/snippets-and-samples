-- name: GetOneURL :one
SELECT * FROM urls
WHERE short_url = ?
LIMIT 1;

-- name: AllURLs :many
SELECT * FROM urls
ORDER BY id DESC;

-- name: CreateURL :one
INSERT INTO urls (short_url, long_url)
VALUES (?, ?)
RETURNING *;

-- name: UpdateVisited :exec
UPDATE urls
SET visited = ?
WHERE short_url = ?;
