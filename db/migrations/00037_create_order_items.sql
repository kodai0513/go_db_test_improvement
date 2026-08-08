-- +goose Up
CREATE TABLE order_items (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price_yen INTEGER NOT NULL
);

-- +goose Down
DROP TABLE order_items;
