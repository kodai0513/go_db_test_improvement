-- +goose Up
CREATE TABLE staff_roles (
    id UUID PRIMARY KEY,
    staff_member_id UUID NOT NULL,
    role_id UUID NOT NULL
);

-- +goose Down
DROP TABLE staff_roles;
