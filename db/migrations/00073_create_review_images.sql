-- +goose Up
CREATE TABLE review_images (
    id UUID PRIMARY KEY,
    review_id UUID NOT NULL,
    url TEXT NOT NULL
);

-- +goose Down
DROP TABLE review_images;
