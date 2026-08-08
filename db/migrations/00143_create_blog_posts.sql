-- +goose Up
CREATE TABLE blog_posts (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    published_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE blog_posts;
