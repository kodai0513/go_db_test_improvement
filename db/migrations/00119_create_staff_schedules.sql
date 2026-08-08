-- +goose Up
CREATE TABLE staff_schedules (
    id UUID PRIMARY KEY,
    staff_member_id UUID NOT NULL,
    work_date TIMESTAMPTZ NOT NULL,
    shift TEXT NOT NULL
);

-- +goose Down
DROP TABLE staff_schedules;
