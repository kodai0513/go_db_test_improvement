-- +goose Up
CREATE TABLE shipment_items (
    id UUID PRIMARY KEY,
    shipment_id UUID NOT NULL,
    order_item_id UUID NOT NULL,
    quantity INTEGER NOT NULL
);

-- +goose Down
DROP TABLE shipment_items;
