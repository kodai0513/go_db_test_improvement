-- +goose Up
CREATE TABLE shipping_addresses (
    id UUID PRIMARY KEY,
    shipment_id UUID NOT NULL,
    postal_code TEXT NOT NULL,
    address TEXT NOT NULL
);

-- +goose Down
DROP TABLE shipping_addresses;
