-- +goose Up
CREATE TABLE review_tag_assignments (
    id UUID PRIMARY KEY,
    review_id UUID NOT NULL,
    review_tag_id UUID NOT NULL
);

-- +goose Down
DROP TABLE review_tag_assignments;
