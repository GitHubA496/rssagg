-- +goose Up

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL ,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    url TEXT UNIQUE NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON Delete cascade
);

-- +goose Down
Drop table posts;