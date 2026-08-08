-- +goose Up
CREATE TABLE vendor_reviews (
    id UUID PRIMARY KEY,
    vendor_id UUID NOT NULL,
    member_id UUID NOT NULL,
    rating INTEGER NOT NULL,
    comment TEXT NOT NULL
);

-- +goose Down
DROP TABLE vendor_reviews;
