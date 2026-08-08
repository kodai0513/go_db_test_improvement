-- +goose Up
CREATE TABLE review_tags (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE review_tags;
