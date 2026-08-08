-- +goose Up
CREATE TABLE menu_items (
    id UUID PRIMARY KEY,
    menu_id UUID NOT NULL,
    label TEXT NOT NULL,
    url TEXT NOT NULL,
    sort_order INTEGER NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE menu_items;
