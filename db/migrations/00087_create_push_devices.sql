-- +goose Up
CREATE TABLE push_devices (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    device_token TEXT NOT NULL UNIQUE,
    registered_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE push_devices;
