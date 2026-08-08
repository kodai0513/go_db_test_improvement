-- +goose Up
CREATE TABLE ticket_assignments (
    id UUID PRIMARY KEY,
    ticket_id UUID NOT NULL,
    staff_member_id UUID NOT NULL,
    assigned_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE ticket_assignments;
