-- +goose Up
CREATE TABLE point_redemptions (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    points INTEGER NOT NULL,
    redeemed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE point_redemptions;
