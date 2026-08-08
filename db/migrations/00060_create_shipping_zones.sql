-- +goose Up
CREATE TABLE shipping_zones (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE shipping_zones;
