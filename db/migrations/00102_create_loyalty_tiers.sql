-- +goose Up
CREATE TABLE loyalty_tiers (
    id UUID PRIMARY KEY,
    loyalty_program_id UUID NOT NULL,
    name TEXT NOT NULL,
    min_points INTEGER NOT NULL
);

-- +goose Down
DROP TABLE loyalty_tiers;
