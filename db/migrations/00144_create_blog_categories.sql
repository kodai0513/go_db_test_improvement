-- +goose Up
CREATE TABLE blog_categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE blog_categories;
