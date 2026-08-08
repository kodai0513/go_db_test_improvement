-- +goose Up
CREATE TABLE reward_redemptions (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    reward_item_id UUID NOT NULL,
    redeemed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE reward_redemptions;
