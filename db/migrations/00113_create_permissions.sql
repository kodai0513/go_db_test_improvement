-- +goose Up
CREATE TABLE permissions (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE permissions;
