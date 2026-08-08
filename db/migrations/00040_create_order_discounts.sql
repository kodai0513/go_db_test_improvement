-- +goose Up
CREATE TABLE order_discounts (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    coupon_id UUID,
    discount_amount_yen INTEGER NOT NULL
);

-- +goose Down
DROP TABLE order_discounts;
