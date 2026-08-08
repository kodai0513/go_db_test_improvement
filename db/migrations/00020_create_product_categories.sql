-- +goose Up
CREATE TABLE product_categories (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    category_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE product_categories;
