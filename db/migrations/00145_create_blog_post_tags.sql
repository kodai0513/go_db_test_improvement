-- +goose Up
CREATE TABLE blog_post_tags (
    id UUID PRIMARY KEY,
    blog_post_id UUID NOT NULL,
    tag TEXT NOT NULL
);

-- +goose Down
DROP TABLE blog_post_tags;
