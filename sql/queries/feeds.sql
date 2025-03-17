-- name: CreateFeed :one
Insert into feeds (id, created_at, updated_at, name, url, user_id)
values ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetFeed :many
SELECT * FROM feeds;

-- name: GetNextFeedToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),updated_at= NOW()
where id = $1
RETURNING *;
