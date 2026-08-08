-- +goose Up
CREATE TABLE pages (
    id UUID PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL,
    body TEXT NOT NULL
);

-- +goose Down
DROP TABLE pages;
