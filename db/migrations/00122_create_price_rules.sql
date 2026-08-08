-- +goose Up
CREATE TABLE price_rules (
    id UUID PRIMARY KEY,
    price_list_id UUID NOT NULL,
    product_id UUID NOT NULL,
    price_yen INTEGER NOT NULL
);

-- +goose Down
DROP TABLE price_rules;
