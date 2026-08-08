-- +goose Up
CREATE TABLE member_referrals (
    id UUID PRIMARY KEY,
    referrer_member_id UUID NOT NULL,
    referred_member_id UUID NOT NULL,
    reward_yen INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE member_referrals;
