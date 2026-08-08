-- name: CreateVendor :one
INSERT INTO vendors (id, name, email, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetVendor :one
SELECT * FROM vendors WHERE id = $1;

-- name: ListVendors :many
SELECT * FROM vendors ORDER BY created_at DESC;

-- name: DeleteVendor :exec
DELETE FROM vendors WHERE id = $1;

-- name: CreateVendorProduct :one
INSERT INTO vendor_products (id, vendor_id, product_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetVendorProduct :one
SELECT * FROM vendor_products WHERE id = $1;

-- name: ListVendorProducts :many
SELECT * FROM vendor_products ORDER BY id DESC;

-- name: ListVendorProductsByVendorID :many
SELECT * FROM vendor_products WHERE vendor_id = $1 ORDER BY id DESC;

-- name: DeleteVendorProduct :exec
DELETE FROM vendor_products WHERE id = $1;

-- name: CreateVendorPayout :one
INSERT INTO vendor_payouts (id, vendor_id, amount_yen, paid_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetVendorPayout :one
SELECT * FROM vendor_payouts WHERE id = $1;

-- name: ListVendorPayouts :many
SELECT * FROM vendor_payouts ORDER BY paid_at DESC;

-- name: ListVendorPayoutsByVendorID :many
SELECT * FROM vendor_payouts WHERE vendor_id = $1 ORDER BY paid_at DESC;

-- name: DeleteVendorPayout :exec
DELETE FROM vendor_payouts WHERE id = $1;

-- name: CreateVendorReview :one
INSERT INTO vendor_reviews (id, vendor_id, member_id, rating, comment)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetVendorReview :one
SELECT * FROM vendor_reviews WHERE id = $1;

-- name: ListVendorReviews :many
SELECT * FROM vendor_reviews ORDER BY id DESC;

-- name: ListVendorReviewsByVendorID :many
SELECT * FROM vendor_reviews WHERE vendor_id = $1 ORDER BY id DESC;

-- name: DeleteVendorReview :exec
DELETE FROM vendor_reviews WHERE id = $1;

-- name: CreateVendorDocument :one
INSERT INTO vendor_documents (id, vendor_id, url, doc_type)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetVendorDocument :one
SELECT * FROM vendor_documents WHERE id = $1;

-- name: ListVendorDocuments :many
SELECT * FROM vendor_documents ORDER BY id DESC;

-- name: ListVendorDocumentsByVendorID :many
SELECT * FROM vendor_documents WHERE vendor_id = $1 ORDER BY id DESC;

-- name: DeleteVendorDocument :exec
DELETE FROM vendor_documents WHERE id = $1;

-- name: CreateVendorCategory :one
INSERT INTO vendor_categories (id, vendor_id, category_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetVendorCategory :one
SELECT * FROM vendor_categories WHERE id = $1;

-- name: ListVendorCategories :many
SELECT * FROM vendor_categories ORDER BY id DESC;

-- name: ListVendorCategoriesByVendorID :many
SELECT * FROM vendor_categories WHERE vendor_id = $1 ORDER BY id DESC;

-- name: DeleteVendorCategory :exec
DELETE FROM vendor_categories WHERE id = $1;

-- name: CreateVendorCommission :one
INSERT INTO vendor_commissions (id, vendor_id, rate_percent)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetVendorCommission :one
SELECT * FROM vendor_commissions WHERE id = $1;

-- name: ListVendorCommissions :many
SELECT * FROM vendor_commissions ORDER BY id DESC;

-- name: ListVendorCommissionsByVendorID :many
SELECT * FROM vendor_commissions WHERE vendor_id = $1 ORDER BY id DESC;

-- name: DeleteVendorCommission :exec
DELETE FROM vendor_commissions WHERE id = $1;

-- name: CreateVendorBankAccount :one
INSERT INTO vendor_bank_accounts (id, vendor_id, account_number)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetVendorBankAccount :one
SELECT * FROM vendor_bank_accounts WHERE id = $1;

-- name: ListVendorBankAccounts :many
SELECT * FROM vendor_bank_accounts ORDER BY id DESC;

-- name: ListVendorBankAccountsByVendorID :many
SELECT * FROM vendor_bank_accounts WHERE vendor_id = $1 ORDER BY id DESC;

-- name: DeleteVendorBankAccount :exec
DELETE FROM vendor_bank_accounts WHERE id = $1;

-- name: CreateVendorContract :one
INSERT INTO vendor_contracts (id, vendor_id, starts_at, ends_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetVendorContract :one
SELECT * FROM vendor_contracts WHERE id = $1;

-- name: ListVendorContracts :many
SELECT * FROM vendor_contracts ORDER BY starts_at DESC;

-- name: ListVendorContractsByVendorID :many
SELECT * FROM vendor_contracts WHERE vendor_id = $1 ORDER BY starts_at DESC;

-- name: DeleteVendorContract :exec
DELETE FROM vendor_contracts WHERE id = $1;

-- name: CreateVendorPerformanceMetric :one
INSERT INTO vendor_performance_metrics (id, vendor_id, metric_name, value, recorded_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetVendorPerformanceMetric :one
SELECT * FROM vendor_performance_metrics WHERE id = $1;

-- name: ListVendorPerformanceMetrics :many
SELECT * FROM vendor_performance_metrics ORDER BY recorded_at DESC;

-- name: ListVendorPerformanceMetricsByVendorID :many
SELECT * FROM vendor_performance_metrics WHERE vendor_id = $1 ORDER BY recorded_at DESC;

-- name: DeleteVendorPerformanceMetric :exec
DELETE FROM vendor_performance_metrics WHERE id = $1;
