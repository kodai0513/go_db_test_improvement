-- name: CreateMember :one
INSERT INTO members (id, name, email, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMember :one
SELECT * FROM members WHERE id = $1;

-- name: ListMembers :many
SELECT * FROM members ORDER BY created_at DESC;

-- name: CreateMemberProfile :one
INSERT INTO member_profiles (id, member_id, birth_date, gender, bio, avatar_url, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetMemberProfile :one
SELECT * FROM member_profiles WHERE id = $1;

-- name: ListMemberProfiles :many
SELECT * FROM member_profiles ORDER BY updated_at DESC;

-- name: ListMemberProfilesByMemberID :many
SELECT * FROM member_profiles WHERE member_id = $1 ORDER BY updated_at DESC;

-- name: DeleteMemberProfile :exec
DELETE FROM member_profiles WHERE id = $1;

-- name: CreateMemberAddress :one
INSERT INTO member_addresses (id, member_id, postal_code, prefecture, city, line1, line2, is_default, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetMemberAddress :one
SELECT * FROM member_addresses WHERE id = $1;

-- name: ListMemberAddresses :many
SELECT * FROM member_addresses ORDER BY created_at DESC;

-- name: ListMemberAddressesByMemberID :many
SELECT * FROM member_addresses WHERE member_id = $1 ORDER BY created_at DESC;

-- name: DeleteMemberAddress :exec
DELETE FROM member_addresses WHERE id = $1;

-- name: CreateMemberCredential :one
INSERT INTO member_credentials (id, member_id, password_hash, last_changed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMemberCredential :one
SELECT * FROM member_credentials WHERE id = $1;

-- name: ListMemberCredentials :many
SELECT * FROM member_credentials ORDER BY last_changed_at DESC;

-- name: ListMemberCredentialsByMemberID :many
SELECT * FROM member_credentials WHERE member_id = $1 ORDER BY last_changed_at DESC;

-- name: DeleteMemberCredential :exec
DELETE FROM member_credentials WHERE id = $1;

-- name: CreateMemberSession :one
INSERT INTO member_sessions (id, member_id, session_token, expires_at, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetMemberSession :one
SELECT * FROM member_sessions WHERE id = $1;

-- name: ListMemberSessions :many
SELECT * FROM member_sessions ORDER BY created_at DESC;

-- name: ListMemberSessionsByMemberID :many
SELECT * FROM member_sessions WHERE member_id = $1 ORDER BY created_at DESC;

-- name: DeleteMemberSession :exec
DELETE FROM member_sessions WHERE id = $1;

-- name: CreateMemberPreference :one
INSERT INTO member_preferences (id, member_id, language, newsletter_opt_in, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetMemberPreference :one
SELECT * FROM member_preferences WHERE id = $1;

-- name: ListMemberPreferences :many
SELECT * FROM member_preferences ORDER BY updated_at DESC;

-- name: ListMemberPreferencesByMemberID :many
SELECT * FROM member_preferences WHERE member_id = $1 ORDER BY updated_at DESC;

-- name: DeleteMemberPreference :exec
DELETE FROM member_preferences WHERE id = $1;

-- name: CreateMemberTag :one
INSERT INTO member_tags (id, name, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetMemberTag :one
SELECT * FROM member_tags WHERE id = $1;

-- name: ListMemberTags :many
SELECT * FROM member_tags ORDER BY created_at DESC;

-- name: DeleteMemberTag :exec
DELETE FROM member_tags WHERE id = $1;

-- name: CreateMemberTagAssignment :one
INSERT INTO member_tag_assignments (id, member_id, member_tag_id, assigned_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMemberTagAssignment :one
SELECT * FROM member_tag_assignments WHERE id = $1;

-- name: ListMemberTagAssignments :many
SELECT * FROM member_tag_assignments ORDER BY assigned_at DESC;

-- name: ListMemberTagAssignmentsByMemberID :many
SELECT * FROM member_tag_assignments WHERE member_id = $1 ORDER BY assigned_at DESC;

-- name: DeleteMemberTagAssignment :exec
DELETE FROM member_tag_assignments WHERE id = $1;

-- name: CreateMemberReferral :one
INSERT INTO member_referrals (id, referrer_member_id, referred_member_id, reward_yen, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetMemberReferral :one
SELECT * FROM member_referrals WHERE id = $1;

-- name: ListMemberReferrals :many
SELECT * FROM member_referrals ORDER BY created_at DESC;

-- name: ListMemberReferralsByReferrerMemberID :many
SELECT * FROM member_referrals WHERE referrer_member_id = $1 ORDER BY created_at DESC;

-- name: DeleteMemberReferral :exec
DELETE FROM member_referrals WHERE id = $1;

-- name: CreateMemberPointAccount :one
INSERT INTO member_point_accounts (id, member_id, balance, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMemberPointAccount :one
SELECT * FROM member_point_accounts WHERE id = $1;

-- name: ListMemberPointAccounts :many
SELECT * FROM member_point_accounts ORDER BY updated_at DESC;

-- name: ListMemberPointAccountsByMemberID :many
SELECT * FROM member_point_accounts WHERE member_id = $1 ORDER BY updated_at DESC;

-- name: DeleteMemberPointAccount :exec
DELETE FROM member_point_accounts WHERE id = $1;
