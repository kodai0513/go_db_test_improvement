-- +goose Up
CREATE TABLE coupon_categories (
    id UUID PRIMARY KEY,
    coupon_id UUID NOT NULL,
    category_id UUID NOT NULL
);

-- +goose Down
DROP TABLE coupon_categories;
