-- +goose Up
CREATE TABLE supplier_products (
    id UUID PRIMARY KEY,
    supplier_id UUID NOT NULL,
    product_id UUID NOT NULL,
    lead_time_days INTEGER NOT NULL
);

-- +goose Down
DROP TABLE supplier_products;
