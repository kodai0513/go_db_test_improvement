-- +goose Up
CREATE TABLE notification_preferences (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    channel_id UUID NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT true
);

-- +goose Down
DROP TABLE notification_preferences;
