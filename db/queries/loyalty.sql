-- name: CreateLoyaltyProgram :one
INSERT INTO loyalty_programs (id, name, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetLoyaltyProgram :one
SELECT * FROM loyalty_programs WHERE id = $1;

-- name: ListLoyaltyPrograms :many
SELECT * FROM loyalty_programs ORDER BY created_at DESC;

-- name: DeleteLoyaltyProgram :exec
DELETE FROM loyalty_programs WHERE id = $1;

-- name: CreateLoyaltyTier :one
INSERT INTO loyalty_tiers (id, loyalty_program_id, name, min_points)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetLoyaltyTier :one
SELECT * FROM loyalty_tiers WHERE id = $1;

-- name: ListLoyaltyTiers :many
SELECT * FROM loyalty_tiers ORDER BY id DESC;

-- name: ListLoyaltyTiersByLoyaltyProgramID :many
SELECT * FROM loyalty_tiers WHERE loyalty_program_id = $1 ORDER BY id DESC;

-- name: DeleteLoyaltyTier :exec
DELETE FROM loyalty_tiers WHERE id = $1;

-- name: CreatePointRule :one
INSERT INTO point_rules (id, loyalty_program_id, action, points)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPointRule :one
SELECT * FROM point_rules WHERE id = $1;

-- name: ListPointRules :many
SELECT * FROM point_rules ORDER BY id DESC;

-- name: ListPointRulesByLoyaltyProgramID :many
SELECT * FROM point_rules WHERE loyalty_program_id = $1 ORDER BY id DESC;

-- name: DeletePointRule :exec
DELETE FROM point_rules WHERE id = $1;

-- name: CreatePointRedemption :one
INSERT INTO point_redemptions (id, member_id, points, redeemed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPointRedemption :one
SELECT * FROM point_redemptions WHERE id = $1;

-- name: ListPointRedemptions :many
SELECT * FROM point_redemptions ORDER BY redeemed_at DESC;

-- name: ListPointRedemptionsByMemberID :many
SELECT * FROM point_redemptions WHERE member_id = $1 ORDER BY redeemed_at DESC;

-- name: DeletePointRedemption :exec
DELETE FROM point_redemptions WHERE id = $1;

-- name: CreateLoyaltyTierHistory :one
INSERT INTO loyalty_tier_histories (id, member_id, loyalty_tier_id, changed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetLoyaltyTierHistory :one
SELECT * FROM loyalty_tier_histories WHERE id = $1;

-- name: ListLoyaltyTierHistories :many
SELECT * FROM loyalty_tier_histories ORDER BY changed_at DESC;

-- name: ListLoyaltyTierHistoriesByMemberID :many
SELECT * FROM loyalty_tier_histories WHERE member_id = $1 ORDER BY changed_at DESC;

-- name: DeleteLoyaltyTierHistory :exec
DELETE FROM loyalty_tier_histories WHERE id = $1;

-- name: CreateRewardItem :one
INSERT INTO reward_items (id, name, points_cost)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetRewardItem :one
SELECT * FROM reward_items WHERE id = $1;

-- name: ListRewardItems :many
SELECT * FROM reward_items ORDER BY id DESC;

-- name: DeleteRewardItem :exec
DELETE FROM reward_items WHERE id = $1;

-- name: CreateRewardRedemption :one
INSERT INTO reward_redemptions (id, member_id, reward_item_id, redeemed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetRewardRedemption :one
SELECT * FROM reward_redemptions WHERE id = $1;

-- name: ListRewardRedemptions :many
SELECT * FROM reward_redemptions ORDER BY redeemed_at DESC;

-- name: ListRewardRedemptionsByMemberID :many
SELECT * FROM reward_redemptions WHERE member_id = $1 ORDER BY redeemed_at DESC;

-- name: DeleteRewardRedemption :exec
DELETE FROM reward_redemptions WHERE id = $1;

-- name: CreateReferralBonus :one
INSERT INTO referral_bonuses (id, member_id, points, granted_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetReferralBonus :one
SELECT * FROM referral_bonuses WHERE id = $1;

-- name: ListReferralBonuses :many
SELECT * FROM referral_bonuses ORDER BY granted_at DESC;

-- name: ListReferralBonusesByMemberID :many
SELECT * FROM referral_bonuses WHERE member_id = $1 ORDER BY granted_at DESC;

-- name: DeleteReferralBonus :exec
DELETE FROM referral_bonuses WHERE id = $1;

-- name: CreateBirthdayReward :one
INSERT INTO birthday_rewards (id, member_id, points, granted_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetBirthdayReward :one
SELECT * FROM birthday_rewards WHERE id = $1;

-- name: ListBirthdayRewards :many
SELECT * FROM birthday_rewards ORDER BY granted_at DESC;

-- name: ListBirthdayRewardsByMemberID :many
SELECT * FROM birthday_rewards WHERE member_id = $1 ORDER BY granted_at DESC;

-- name: DeleteBirthdayReward :exec
DELETE FROM birthday_rewards WHERE id = $1;

-- name: CreateLoyaltyPointExpiration :one
INSERT INTO loyalty_point_expirations (id, member_id, points, expires_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetLoyaltyPointExpiration :one
SELECT * FROM loyalty_point_expirations WHERE id = $1;

-- name: ListLoyaltyPointExpirations :many
SELECT * FROM loyalty_point_expirations ORDER BY expires_at DESC;

-- name: ListLoyaltyPointExpirationsByMemberID :many
SELECT * FROM loyalty_point_expirations WHERE member_id = $1 ORDER BY expires_at DESC;

-- name: DeleteLoyaltyPointExpiration :exec
DELETE FROM loyalty_point_expirations WHERE id = $1;
