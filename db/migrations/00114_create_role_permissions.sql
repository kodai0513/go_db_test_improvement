-- +goose Up
CREATE TABLE role_permissions (
    id UUID PRIMARY KEY,
    role_id UUID NOT NULL,
    permission_id UUID NOT NULL
);

-- +goose Down
DROP TABLE role_permissions;
