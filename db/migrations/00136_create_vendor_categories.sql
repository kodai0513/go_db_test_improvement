-- +goose Up
CREATE TABLE vendor_categories (
    id UUID PRIMARY KEY,
    vendor_id UUID NOT NULL,
    category_id UUID NOT NULL
);

-- +goose Down
DROP TABLE vendor_categories;
