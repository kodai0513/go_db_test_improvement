-- +goose Up
CREATE TABLE loyalty_tier_histories (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    loyalty_tier_id UUID NOT NULL,
    changed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE loyalty_tier_histories;
