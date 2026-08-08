-- +goose Up
CREATE TABLE product_tag_assignments (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    product_tag_id UUID NOT NULL
);

-- +goose Down
DROP TABLE product_tag_assignments;
