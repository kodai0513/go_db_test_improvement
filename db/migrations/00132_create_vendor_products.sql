-- +goose Up
CREATE TABLE vendor_products (
    id UUID PRIMARY KEY,
    vendor_id UUID NOT NULL,
    product_id UUID NOT NULL
);

-- +goose Down
DROP TABLE vendor_products;
