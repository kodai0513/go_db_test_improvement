-- +goose Up
CREATE TABLE vendor_documents (
    id UUID PRIMARY KEY,
    vendor_id UUID NOT NULL,
    url TEXT NOT NULL,
    doc_type TEXT NOT NULL
);

-- +goose Down
DROP TABLE vendor_documents;
