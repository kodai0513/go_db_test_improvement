-- +goose Up
CREATE TABLE installment_plans (
    id UUID PRIMARY KEY,
    payment_id UUID NOT NULL,
    number_of_installments INTEGER NOT NULL
);

-- +goose Down
DROP TABLE installment_plans;
