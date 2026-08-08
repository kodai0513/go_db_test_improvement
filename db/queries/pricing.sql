-- name: CreatePriceList :one
INSERT INTO price_lists (id, name, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPriceList :one
SELECT * FROM price_lists WHERE id = $1;

-- name: ListPriceLists :many
SELECT * FROM price_lists ORDER BY created_at DESC;

-- name: DeletePriceList :exec
DELETE FROM price_lists WHERE id = $1;

-- name: CreatePriceRule :one
INSERT INTO price_rules (id, price_list_id, product_id, price_yen)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPriceRule :one
SELECT * FROM price_rules WHERE id = $1;

-- name: ListPriceRules :many
SELECT * FROM price_rules ORDER BY id DESC;

-- name: ListPriceRulesByPriceListID :many
SELECT * FROM price_rules WHERE price_list_id = $1 ORDER BY id DESC;

-- name: DeletePriceRule :exec
DELETE FROM price_rules WHERE id = $1;

-- name: CreateTaxRate :one
INSERT INTO tax_rates (id, name, rate)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTaxRate :one
SELECT * FROM tax_rates WHERE id = $1;

-- name: ListTaxRates :many
SELECT * FROM tax_rates ORDER BY id DESC;

-- name: DeleteTaxRate :exec
DELETE FROM tax_rates WHERE id = $1;

-- name: CreateTaxRule :one
INSERT INTO tax_rules (id, tax_rate_id, region)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTaxRule :one
SELECT * FROM tax_rules WHERE id = $1;

-- name: ListTaxRules :many
SELECT * FROM tax_rules ORDER BY id DESC;

-- name: ListTaxRulesByTaxRateID :many
SELECT * FROM tax_rules WHERE tax_rate_id = $1 ORDER BY id DESC;

-- name: DeleteTaxRule :exec
DELETE FROM tax_rules WHERE id = $1;

-- name: CreateCurrency :one
INSERT INTO currencies (id, code)
VALUES ($1, $2)
RETURNING *;

-- name: GetCurrency :one
SELECT * FROM currencies WHERE id = $1;

-- name: ListCurrencies :many
SELECT * FROM currencies ORDER BY id DESC;

-- name: DeleteCurrency :exec
DELETE FROM currencies WHERE id = $1;

-- name: CreateCurrencyRate :one
INSERT INTO currency_rates (id, currency_id, rate_to_jpy, effective_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCurrencyRate :one
SELECT * FROM currency_rates WHERE id = $1;

-- name: ListCurrencyRates :many
SELECT * FROM currency_rates ORDER BY effective_at DESC;

-- name: ListCurrencyRatesByCurrencyID :many
SELECT * FROM currency_rates WHERE currency_id = $1 ORDER BY effective_at DESC;

-- name: DeleteCurrencyRate :exec
DELETE FROM currency_rates WHERE id = $1;

-- name: CreateDiscountTier :one
INSERT INTO discount_tiers (id, min_quantity, discount_percent)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetDiscountTier :one
SELECT * FROM discount_tiers WHERE id = $1;

-- name: ListDiscountTiers :many
SELECT * FROM discount_tiers ORDER BY id DESC;

-- name: DeleteDiscountTier :exec
DELETE FROM discount_tiers WHERE id = $1;

-- name: CreatePriceHistory :one
INSERT INTO price_histories (id, product_id, price_yen, changed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPriceHistory :one
SELECT * FROM price_histories WHERE id = $1;

-- name: ListPriceHistories :many
SELECT * FROM price_histories ORDER BY changed_at DESC;

-- name: ListPriceHistoriesByProductID :many
SELECT * FROM price_histories WHERE product_id = $1 ORDER BY changed_at DESC;

-- name: DeletePriceHistory :exec
DELETE FROM price_histories WHERE id = $1;

-- name: CreateBundleOffer :one
INSERT INTO bundle_offers (id, name, price_yen)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetBundleOffer :one
SELECT * FROM bundle_offers WHERE id = $1;

-- name: ListBundleOffers :many
SELECT * FROM bundle_offers ORDER BY id DESC;

-- name: DeleteBundleOffer :exec
DELETE FROM bundle_offers WHERE id = $1;

-- name: CreateBundleOfferItem :one
INSERT INTO bundle_offer_items (id, bundle_offer_id, product_id, quantity)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetBundleOfferItem :one
SELECT * FROM bundle_offer_items WHERE id = $1;

-- name: ListBundleOfferItems :many
SELECT * FROM bundle_offer_items ORDER BY id DESC;

-- name: ListBundleOfferItemsByBundleOfferID :many
SELECT * FROM bundle_offer_items WHERE bundle_offer_id = $1 ORDER BY id DESC;

-- name: DeleteBundleOfferItem :exec
DELETE FROM bundle_offer_items WHERE id = $1;
