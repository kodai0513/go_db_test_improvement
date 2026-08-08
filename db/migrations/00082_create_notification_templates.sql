-- +goose Up
CREATE TABLE notification_templates (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    body TEXT NOT NULL
);

-- +goose Down
DROP TABLE notification_templates;
