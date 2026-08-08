-- +goose Up
CREATE TABLE coupon_stacking_rules (
    id UUID PRIMARY KEY,
    coupon_id UUID NOT NULL,
    stackable_with_coupon_id UUID NOT NULL
);

-- +goose Down
DROP TABLE coupon_stacking_rules;
