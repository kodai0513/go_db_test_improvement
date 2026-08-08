-- +goose Up
CREATE TABLE notification_subscriptions (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    topic TEXT NOT NULL,
    subscribed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE notification_subscriptions;
