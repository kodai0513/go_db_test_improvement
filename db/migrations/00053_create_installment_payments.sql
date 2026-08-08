-- +goose Up
CREATE TABLE installment_payments (
    id UUID PRIMARY KEY,
    installment_plan_id UUID NOT NULL,
    amount_yen INTEGER NOT NULL,
    due_at TIMESTAMPTZ NOT NULL,
    paid_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE installment_payments;
