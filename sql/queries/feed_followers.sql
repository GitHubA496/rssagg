-- name: CreateFeedFollower :one
INSERT INTO feed_followers (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedFollowers :many
SELECT * FROM feed_followers
WHERE user_id = $1;

-- name: DeleteFeedFollower :exec
DELETE FROM feed_followers
WHERE id=$1 and user_id=$2;