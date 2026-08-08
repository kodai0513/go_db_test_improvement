-- +goose Up
CREATE TABLE bundle_offer_items (
    id UUID PRIMARY KEY,
    bundle_offer_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INTEGER NOT NULL
);

-- +goose Down
DROP TABLE bundle_offer_items;
