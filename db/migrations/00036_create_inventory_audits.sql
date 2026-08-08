-- +goose Up
CREATE TABLE inventory_audits (
    id UUID PRIMARY KEY,
    warehouse_id UUID NOT NULL,
    performed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    result TEXT NOT NULL
);

-- +goose Down
DROP TABLE inventory_audits;
