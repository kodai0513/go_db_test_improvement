-- +goose Up
CREATE TABLE discount_tiers (
    id UUID PRIMARY KEY,
    min_quantity INTEGER NOT NULL,
    discount_percent INTEGER NOT NULL
);

-- +goose Down
DROP TABLE discount_tiers;
