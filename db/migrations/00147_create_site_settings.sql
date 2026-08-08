-- +goose Up
CREATE TABLE site_settings (
    id UUID PRIMARY KEY,
    setting_key TEXT NOT NULL UNIQUE,
    setting_value TEXT NOT NULL
);

-- +goose Down
DROP TABLE site_settings;
