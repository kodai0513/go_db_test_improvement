-- +goose Up
CREATE TABLE discount_rules (
    id UUID PRIMARY KEY,
    coupon_id UUID NOT NULL,
    min_amount_yen INTEGER NOT NULL
);

-- +goose Down
DROP TABLE discount_rules;
