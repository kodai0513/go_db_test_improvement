-- +goose Up
CREATE TABLE purchase_order_items (
    id UUID PRIMARY KEY,
    purchase_order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price_yen INTEGER NOT NULL
);

-- +goose Down
DROP TABLE purchase_order_items;
