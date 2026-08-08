-- name: CreatePayment :one
INSERT INTO payments (id, order_id, amount_yen, method, paid_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments WHERE id = $1;

-- name: ListPayments :many
SELECT * FROM payments ORDER BY paid_at DESC;

-- name: CreatePaymentMethod :one
INSERT INTO payment_methods (id, member_id, type, is_default, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPaymentMethod :one
SELECT * FROM payment_methods WHERE id = $1;

-- name: ListPaymentMethods :many
SELECT * FROM payment_methods ORDER BY created_at DESC;

-- name: ListPaymentMethodsByMemberID :many
SELECT * FROM payment_methods WHERE member_id = $1 ORDER BY created_at DESC;

-- name: DeletePaymentMethod :exec
DELETE FROM payment_methods WHERE id = $1;

-- name: CreatePaymentTransaction :one
INSERT INTO payment_transactions (id, payment_id, status, processed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPaymentTransaction :one
SELECT * FROM payment_transactions WHERE id = $1;

-- name: ListPaymentTransactions :many
SELECT * FROM payment_transactions ORDER BY processed_at DESC;

-- name: ListPaymentTransactionsByPaymentID :many
SELECT * FROM payment_transactions WHERE payment_id = $1 ORDER BY processed_at DESC;

-- name: DeletePaymentTransaction :exec
DELETE FROM payment_transactions WHERE id = $1;

-- name: CreateRefund :one
INSERT INTO refunds (id, payment_id, amount_yen, reason, refunded_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetRefund :one
SELECT * FROM refunds WHERE id = $1;

-- name: ListRefunds :many
SELECT * FROM refunds ORDER BY refunded_at DESC;

-- name: ListRefundsByPaymentID :many
SELECT * FROM refunds WHERE payment_id = $1 ORDER BY refunded_at DESC;

-- name: DeleteRefund :exec
DELETE FROM refunds WHERE id = $1;

-- name: CreateInvoice :one
INSERT INTO invoices (id, order_id, amount_yen, issued_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetInvoice :one
SELECT * FROM invoices WHERE id = $1;

-- name: ListInvoices :many
SELECT * FROM invoices ORDER BY issued_at DESC;

-- name: ListInvoicesByOrderID :many
SELECT * FROM invoices WHERE order_id = $1 ORDER BY issued_at DESC;

-- name: DeleteInvoice :exec
DELETE FROM invoices WHERE id = $1;

-- name: CreatePaymentProvider :one
INSERT INTO payment_providers (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetPaymentProvider :one
SELECT * FROM payment_providers WHERE id = $1;

-- name: ListPaymentProviders :many
SELECT * FROM payment_providers ORDER BY id DESC;

-- name: DeletePaymentProvider :exec
DELETE FROM payment_providers WHERE id = $1;

-- name: CreatePaymentProviderAccount :one
INSERT INTO payment_provider_accounts (id, payment_provider_id, member_id, external_account_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPaymentProviderAccount :one
SELECT * FROM payment_provider_accounts WHERE id = $1;

-- name: ListPaymentProviderAccounts :many
SELECT * FROM payment_provider_accounts ORDER BY id DESC;

-- name: ListPaymentProviderAccountsByPaymentProviderID :many
SELECT * FROM payment_provider_accounts WHERE payment_provider_id = $1 ORDER BY id DESC;

-- name: ListPaymentProviderAccountsByMemberID :many
SELECT * FROM payment_provider_accounts WHERE member_id = $1 ORDER BY id DESC;

-- name: DeletePaymentProviderAccount :exec
DELETE FROM payment_provider_accounts WHERE id = $1;

-- name: CreateInstallmentPlan :one
INSERT INTO installment_plans (id, payment_id, number_of_installments)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetInstallmentPlan :one
SELECT * FROM installment_plans WHERE id = $1;

-- name: ListInstallmentPlans :many
SELECT * FROM installment_plans ORDER BY id DESC;

-- name: ListInstallmentPlansByPaymentID :many
SELECT * FROM installment_plans WHERE payment_id = $1 ORDER BY id DESC;

-- name: DeleteInstallmentPlan :exec
DELETE FROM installment_plans WHERE id = $1;

-- name: CreateInstallmentPayment :one
INSERT INTO installment_payments (id, installment_plan_id, amount_yen, due_at, paid_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetInstallmentPayment :one
SELECT * FROM installment_payments WHERE id = $1;

-- name: ListInstallmentPayments :many
SELECT * FROM installment_payments ORDER BY due_at DESC;

-- name: ListInstallmentPaymentsByInstallmentPlanID :many
SELECT * FROM installment_payments WHERE installment_plan_id = $1 ORDER BY due_at DESC;

-- name: DeleteInstallmentPayment :exec
DELETE FROM installment_payments WHERE id = $1;

-- name: CreatePaymentDispute :one
INSERT INTO payment_disputes (id, payment_id, reason, status, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPaymentDispute :one
SELECT * FROM payment_disputes WHERE id = $1;

-- name: ListPaymentDisputes :many
SELECT * FROM payment_disputes ORDER BY created_at DESC;

-- name: ListPaymentDisputesByPaymentID :many
SELECT * FROM payment_disputes WHERE payment_id = $1 ORDER BY created_at DESC;

-- name: DeletePaymentDispute :exec
DELETE FROM payment_disputes WHERE id = $1;
