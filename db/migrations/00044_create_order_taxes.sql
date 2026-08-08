-- +goose Up
CREATE TABLE order_taxes (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    tax_amount_yen INTEGER NOT NULL,
    tax_rate INTEGER NOT NULL
);

-- +goose Down
DROP TABLE order_taxes;
