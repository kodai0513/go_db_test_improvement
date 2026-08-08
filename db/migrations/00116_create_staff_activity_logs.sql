-- +goose Up
CREATE TABLE staff_activity_logs (
    id UUID PRIMARY KEY,
    staff_member_id UUID NOT NULL,
    action TEXT NOT NULL,
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE staff_activity_logs;
