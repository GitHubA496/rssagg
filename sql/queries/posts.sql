-- name: CreatePost :one
Insert into posts (id,
    created_at,
    updated_at,
    title,
    description,
    published_at,
    url,
    feed_id
)
values ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts
Join feed_followers on feed_followers.id = posts.feed_id
where feed_followers.user_id= $1
order by posts.published_at desc 
LIMIT $2; 