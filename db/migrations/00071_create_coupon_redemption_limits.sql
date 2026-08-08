-- +goose Up
CREATE TABLE coupon_redemption_limits (
    id UUID PRIMARY KEY,
    coupon_id UUID NOT NULL,
    max_redemptions INTEGER NOT NULL
);

-- +goose Down
DROP TABLE coupon_redemption_limits;
