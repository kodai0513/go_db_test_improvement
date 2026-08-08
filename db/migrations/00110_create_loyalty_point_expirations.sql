-- +goose Up
CREATE TABLE loyalty_point_expirations (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    points INTEGER NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL
);

-- +goose Down
DROP TABLE loyalty_point_expirations;
