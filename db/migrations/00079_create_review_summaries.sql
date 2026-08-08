-- +goose Up
CREATE TABLE review_summaries (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL UNIQUE,
    average_rating INTEGER NOT NULL,
    review_count INTEGER NOT NULL
);

-- +goose Down
DROP TABLE review_summaries;
