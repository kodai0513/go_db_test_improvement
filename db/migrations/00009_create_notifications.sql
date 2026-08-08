-- +goose Up
CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    message TEXT NOT NULL,
    is_read BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE notifications;
