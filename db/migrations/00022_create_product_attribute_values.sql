-- +goose Up
CREATE TABLE product_attribute_values (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    product_attribute_id UUID NOT NULL,
    value TEXT NOT NULL
);

-- +goose Down
DROP TABLE product_attribute_values;
