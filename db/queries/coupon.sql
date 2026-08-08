-- name: CreateCoupon :one
INSERT INTO coupons (id, code, discount_rate, expires_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCoupon :one
SELECT * FROM coupons WHERE id = $1;

-- name: ListCoupons :many
SELECT * FROM coupons ORDER BY expires_at;

-- name: CreateCouponUsage :one
INSERT INTO coupon_usages (id, coupon_id, order_id, used_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCouponUsage :one
SELECT * FROM coupon_usages WHERE id = $1;

-- name: ListCouponUsages :many
SELECT * FROM coupon_usages ORDER BY used_at DESC;

-- name: ListCouponUsagesByCouponID :many
SELECT * FROM coupon_usages WHERE coupon_id = $1 ORDER BY used_at DESC;

-- name: DeleteCouponUsage :exec
DELETE FROM coupon_usages WHERE id = $1;

-- name: CreateCampaign :one
INSERT INTO campaigns (id, name, starts_at, ends_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCampaign :one
SELECT * FROM campaigns WHERE id = $1;

-- name: ListCampaigns :many
SELECT * FROM campaigns ORDER BY id DESC;

-- name: DeleteCampaign :exec
DELETE FROM campaigns WHERE id = $1;

-- name: CreateCampaignTarget :one
INSERT INTO campaign_targets (id, campaign_id, member_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCampaignTarget :one
SELECT * FROM campaign_targets WHERE id = $1;

-- name: ListCampaignTargets :many
SELECT * FROM campaign_targets ORDER BY id DESC;

-- name: ListCampaignTargetsByCampaignID :many
SELECT * FROM campaign_targets WHERE campaign_id = $1 ORDER BY id DESC;

-- name: DeleteCampaignTarget :exec
DELETE FROM campaign_targets WHERE id = $1;

-- name: CreateDiscountRule :one
INSERT INTO discount_rules (id, coupon_id, min_amount_yen)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetDiscountRule :one
SELECT * FROM discount_rules WHERE id = $1;

-- name: ListDiscountRules :many
SELECT * FROM discount_rules ORDER BY id DESC;

-- name: ListDiscountRulesByCouponID :many
SELECT * FROM discount_rules WHERE coupon_id = $1 ORDER BY id DESC;

-- name: DeleteDiscountRule :exec
DELETE FROM discount_rules WHERE id = $1;

-- name: CreateCouponCategory :one
INSERT INTO coupon_categories (id, coupon_id, category_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCouponCategory :one
SELECT * FROM coupon_categories WHERE id = $1;

-- name: ListCouponCategories :many
SELECT * FROM coupon_categories ORDER BY id DESC;

-- name: ListCouponCategoriesByCouponID :many
SELECT * FROM coupon_categories WHERE coupon_id = $1 ORDER BY id DESC;

-- name: DeleteCouponCategory :exec
DELETE FROM coupon_categories WHERE id = $1;

-- name: CreateCouponProduct :one
INSERT INTO coupon_products (id, coupon_id, product_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCouponProduct :one
SELECT * FROM coupon_products WHERE id = $1;

-- name: ListCouponProducts :many
SELECT * FROM coupon_products ORDER BY id DESC;

-- name: ListCouponProductsByCouponID :many
SELECT * FROM coupon_products WHERE coupon_id = $1 ORDER BY id DESC;

-- name: DeleteCouponProduct :exec
DELETE FROM coupon_products WHERE id = $1;

-- name: CreatePromoCode :one
INSERT INTO promo_codes (id, code, campaign_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPromoCode :one
SELECT * FROM promo_codes WHERE id = $1;

-- name: ListPromoCodes :many
SELECT * FROM promo_codes ORDER BY id DESC;

-- name: ListPromoCodesByCampaignID :many
SELECT * FROM promo_codes WHERE campaign_id = $1 ORDER BY id DESC;

-- name: DeletePromoCode :exec
DELETE FROM promo_codes WHERE id = $1;

-- name: CreateCouponRedemptionLimit :one
INSERT INTO coupon_redemption_limits (id, coupon_id, max_redemptions)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCouponRedemptionLimit :one
SELECT * FROM coupon_redemption_limits WHERE id = $1;

-- name: ListCouponRedemptionLimits :many
SELECT * FROM coupon_redemption_limits ORDER BY id DESC;

-- name: ListCouponRedemptionLimitsByCouponID :many
SELECT * FROM coupon_redemption_limits WHERE coupon_id = $1 ORDER BY id DESC;

-- name: DeleteCouponRedemptionLimit :exec
DELETE FROM coupon_redemption_limits WHERE id = $1;

-- name: CreateCouponStackingRule :one
INSERT INTO coupon_stacking_rules (id, coupon_id, stackable_with_coupon_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCouponStackingRule :one
SELECT * FROM coupon_stacking_rules WHERE id = $1;

-- name: ListCouponStackingRules :many
SELECT * FROM coupon_stacking_rules ORDER BY id DESC;

-- name: ListCouponStackingRulesByCouponID :many
SELECT * FROM coupon_stacking_rules WHERE coupon_id = $1 ORDER BY id DESC;

-- name: DeleteCouponStackingRule :exec
DELETE FROM coupon_stacking_rules WHERE id = $1;
