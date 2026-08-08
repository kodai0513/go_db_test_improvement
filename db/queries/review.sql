-- name: CreateReview :one
INSERT INTO reviews (id, product_id, member_id, rating, comment, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetReview :one
SELECT * FROM reviews WHERE id = $1;

-- name: ListReviews :many
SELECT * FROM reviews ORDER BY created_at DESC;

-- name: CreateReviewImage :one
INSERT INTO review_images (id, review_id, url)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetReviewImage :one
SELECT * FROM review_images WHERE id = $1;

-- name: ListReviewImages :many
SELECT * FROM review_images ORDER BY id DESC;

-- name: ListReviewImagesByReviewID :many
SELECT * FROM review_images WHERE review_id = $1 ORDER BY id DESC;

-- name: DeleteReviewImage :exec
DELETE FROM review_images WHERE id = $1;

-- name: CreateReviewVote :one
INSERT INTO review_votes (id, review_id, member_id, is_helpful)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetReviewVote :one
SELECT * FROM review_votes WHERE id = $1;

-- name: ListReviewVotes :many
SELECT * FROM review_votes ORDER BY id DESC;

-- name: ListReviewVotesByReviewID :many
SELECT * FROM review_votes WHERE review_id = $1 ORDER BY id DESC;

-- name: DeleteReviewVote :exec
DELETE FROM review_votes WHERE id = $1;

-- name: CreateReviewReply :one
INSERT INTO review_replies (id, review_id, member_id, comment, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetReviewReply :one
SELECT * FROM review_replies WHERE id = $1;

-- name: ListReviewReplies :many
SELECT * FROM review_replies ORDER BY created_at DESC;

-- name: ListReviewRepliesByReviewID :many
SELECT * FROM review_replies WHERE review_id = $1 ORDER BY created_at DESC;

-- name: DeleteReviewReply :exec
DELETE FROM review_replies WHERE id = $1;

-- name: CreateReviewReport :one
INSERT INTO review_reports (id, review_id, reporter_member_id, reason, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetReviewReport :one
SELECT * FROM review_reports WHERE id = $1;

-- name: ListReviewReports :many
SELECT * FROM review_reports ORDER BY created_at DESC;

-- name: ListReviewReportsByReviewID :many
SELECT * FROM review_reports WHERE review_id = $1 ORDER BY created_at DESC;

-- name: DeleteReviewReport :exec
DELETE FROM review_reports WHERE id = $1;

-- name: CreateReviewTag :one
INSERT INTO review_tags (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetReviewTag :one
SELECT * FROM review_tags WHERE id = $1;

-- name: ListReviewTags :many
SELECT * FROM review_tags ORDER BY id DESC;

-- name: DeleteReviewTag :exec
DELETE FROM review_tags WHERE id = $1;

-- name: CreateReviewTagAssignment :one
INSERT INTO review_tag_assignments (id, review_id, review_tag_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetReviewTagAssignment :one
SELECT * FROM review_tag_assignments WHERE id = $1;

-- name: ListReviewTagAssignments :many
SELECT * FROM review_tag_assignments ORDER BY id DESC;

-- name: ListReviewTagAssignmentsByReviewID :many
SELECT * FROM review_tag_assignments WHERE review_id = $1 ORDER BY id DESC;

-- name: DeleteReviewTagAssignment :exec
DELETE FROM review_tag_assignments WHERE id = $1;

-- name: CreateReviewSummary :one
INSERT INTO review_summaries (id, product_id, average_rating, review_count)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetReviewSummary :one
SELECT * FROM review_summaries WHERE id = $1;

-- name: ListReviewSummaries :many
SELECT * FROM review_summaries ORDER BY id DESC;

-- name: ListReviewSummariesByProductID :many
SELECT * FROM review_summaries WHERE product_id = $1 ORDER BY id DESC;

-- name: DeleteReviewSummary :exec
DELETE FROM review_summaries WHERE id = $1;

-- name: CreateReviewInvitation :one
INSERT INTO review_invitations (id, order_id, member_id, sent_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetReviewInvitation :one
SELECT * FROM review_invitations WHERE id = $1;

-- name: ListReviewInvitations :many
SELECT * FROM review_invitations ORDER BY sent_at DESC;

-- name: ListReviewInvitationsByOrderID :many
SELECT * FROM review_invitations WHERE order_id = $1 ORDER BY sent_at DESC;

-- name: DeleteReviewInvitation :exec
DELETE FROM review_invitations WHERE id = $1;

-- name: CreateReviewModerationLog :one
INSERT INTO review_moderation_logs (id, review_id, action, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetReviewModerationLog :one
SELECT * FROM review_moderation_logs WHERE id = $1;

-- name: ListReviewModerationLogs :many
SELECT * FROM review_moderation_logs ORDER BY created_at DESC;

-- name: ListReviewModerationLogsByReviewID :many
SELECT * FROM review_moderation_logs WHERE review_id = $1 ORDER BY created_at DESC;

-- name: DeleteReviewModerationLog :exec
DELETE FROM review_moderation_logs WHERE id = $1;
