-- +goose Up
CREATE TABLE reward_items (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    points_cost INTEGER NOT NULL
);

-- +goose Down
DROP TABLE reward_items;
