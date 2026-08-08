-- +goose Up
CREATE TABLE staff_departments (
    id UUID PRIMARY KEY,
    staff_member_id UUID NOT NULL,
    department_id UUID NOT NULL
);

-- +goose Down
DROP TABLE staff_departments;
