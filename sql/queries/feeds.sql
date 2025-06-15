-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id, last_fetched_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: GetFeedByUrl :one
SELECT *
FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT *
FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = NOW(), last_fetched_at = NOW()
WHERE id = $1;

-- name: GetNextFeedFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;