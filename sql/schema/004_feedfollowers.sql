-- +goose Up

CREATE TABLE feed_followers (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL ,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON Delete cascade,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON Delete cascade,
    UNIQUE(user_id, feed_id)
);

-- +goose Down
Drop table feed_followers;