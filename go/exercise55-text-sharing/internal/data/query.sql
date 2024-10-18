-- -- name: GetOneURL :one
-- SELECT * FROM urls
-- WHERE short_url = ?
-- LIMIT 1;

-- name: GetAllSnippets :many
SELECT * FROM snippets
ORDER BY created_at DESC;

-- -- name: CreateURL :one
-- INSERT INTO urls (short_url, long_url)
-- VALUES (?, ?)
-- RETURNING *;

-- -- name: UpdateVisited :exec
-- UPDATE urls
-- SET visited = ?
-- WHERE short_url = ?;
