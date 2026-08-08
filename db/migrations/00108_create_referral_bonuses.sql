-- +goose Up
CREATE TABLE referral_bonuses (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    points INTEGER NOT NULL,
    granted_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE referral_bonuses;
