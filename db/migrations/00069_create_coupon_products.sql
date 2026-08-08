-- +goose Up
CREATE TABLE coupon_products (
    id UUID PRIMARY KEY,
    coupon_id UUID NOT NULL,
    product_id UUID NOT NULL
);

-- +goose Down
DROP TABLE coupon_products;
