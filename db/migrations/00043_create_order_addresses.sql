-- +goose Up
CREATE TABLE order_addresses (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    postal_code TEXT NOT NULL,
    address TEXT NOT NULL
);

-- +goose Down
DROP TABLE order_addresses;
