-- +goose Up
CREATE TABLE return_items (
    id UUID PRIMARY KEY,
    return_request_id UUID NOT NULL,
    order_item_id UUID NOT NULL,
    quantity INTEGER NOT NULL
);

-- +goose Down
DROP TABLE return_items;
