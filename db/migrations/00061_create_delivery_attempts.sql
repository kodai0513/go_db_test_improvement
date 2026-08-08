-- +goose Up
CREATE TABLE delivery_attempts (
    id UUID PRIMARY KEY,
    shipment_id UUID NOT NULL,
    attempted_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    succeeded BOOLEAN NOT NULL
);

-- +goose Down
DROP TABLE delivery_attempts;
