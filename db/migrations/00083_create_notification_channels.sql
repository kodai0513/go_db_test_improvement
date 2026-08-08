-- +goose Up
CREATE TABLE notification_channels (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE notification_channels;
