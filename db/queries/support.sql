-- name: CreateTicket :one
INSERT INTO tickets (id, member_id, subject, status, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTicket :one
SELECT * FROM tickets WHERE id = $1;

-- name: ListTickets :many
SELECT * FROM tickets ORDER BY created_at DESC;

-- name: ListTicketsByMemberID :many
SELECT * FROM tickets WHERE member_id = $1 ORDER BY created_at DESC;

-- name: DeleteTicket :exec
DELETE FROM tickets WHERE id = $1;

-- name: CreateTicketMessage :one
INSERT INTO ticket_messages (id, ticket_id, sender_member_id, body, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTicketMessage :one
SELECT * FROM ticket_messages WHERE id = $1;

-- name: ListTicketMessages :many
SELECT * FROM ticket_messages ORDER BY created_at DESC;

-- name: ListTicketMessagesByTicketID :many
SELECT * FROM ticket_messages WHERE ticket_id = $1 ORDER BY created_at DESC;

-- name: DeleteTicketMessage :exec
DELETE FROM ticket_messages WHERE id = $1;

-- name: CreateTicketAttachment :one
INSERT INTO ticket_attachments (id, ticket_message_id, url)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTicketAttachment :one
SELECT * FROM ticket_attachments WHERE id = $1;

-- name: ListTicketAttachments :many
SELECT * FROM ticket_attachments ORDER BY id DESC;

-- name: ListTicketAttachmentsByTicketMessageID :many
SELECT * FROM ticket_attachments WHERE ticket_message_id = $1 ORDER BY id DESC;

-- name: DeleteTicketAttachment :exec
DELETE FROM ticket_attachments WHERE id = $1;

-- name: CreateTicketCategory :one
INSERT INTO ticket_categories (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetTicketCategory :one
SELECT * FROM ticket_categories WHERE id = $1;

-- name: ListTicketCategories :many
SELECT * FROM ticket_categories ORDER BY id DESC;

-- name: DeleteTicketCategory :exec
DELETE FROM ticket_categories WHERE id = $1;

-- name: CreateTicketAssignment :one
INSERT INTO ticket_assignments (id, ticket_id, staff_member_id, assigned_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetTicketAssignment :one
SELECT * FROM ticket_assignments WHERE id = $1;

-- name: ListTicketAssignments :many
SELECT * FROM ticket_assignments ORDER BY assigned_at DESC;

-- name: ListTicketAssignmentsByTicketID :many
SELECT * FROM ticket_assignments WHERE ticket_id = $1 ORDER BY assigned_at DESC;

-- name: DeleteTicketAssignment :exec
DELETE FROM ticket_assignments WHERE id = $1;

-- name: CreateTicketStatusHistory :one
INSERT INTO ticket_status_histories (id, ticket_id, status, changed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetTicketStatusHistory :one
SELECT * FROM ticket_status_histories WHERE id = $1;

-- name: ListTicketStatusHistories :many
SELECT * FROM ticket_status_histories ORDER BY changed_at DESC;

-- name: ListTicketStatusHistoriesByTicketID :many
SELECT * FROM ticket_status_histories WHERE ticket_id = $1 ORDER BY changed_at DESC;

-- name: DeleteTicketStatusHistory :exec
DELETE FROM ticket_status_histories WHERE id = $1;

-- name: CreateFaq :one
INSERT INTO faqs (id, question, answer, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetFaq :one
SELECT * FROM faqs WHERE id = $1;

-- name: ListFaqs :many
SELECT * FROM faqs ORDER BY created_at DESC;

-- name: DeleteFaq :exec
DELETE FROM faqs WHERE id = $1;

-- name: CreateFaqCategory :one
INSERT INTO faq_categories (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetFaqCategory :one
SELECT * FROM faq_categories WHERE id = $1;

-- name: ListFaqCategories :many
SELECT * FROM faq_categories ORDER BY id DESC;

-- name: DeleteFaqCategory :exec
DELETE FROM faq_categories WHERE id = $1;

-- name: CreateSupportAgent :one
INSERT INTO support_agents (id, name, email)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetSupportAgent :one
SELECT * FROM support_agents WHERE id = $1;

-- name: ListSupportAgents :many
SELECT * FROM support_agents ORDER BY id DESC;

-- name: DeleteSupportAgent :exec
DELETE FROM support_agents WHERE id = $1;

-- name: CreateSatisfactionSurvey :one
INSERT INTO satisfaction_surveys (id, ticket_id, score, comment, submitted_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSatisfactionSurvey :one
SELECT * FROM satisfaction_surveys WHERE id = $1;

-- name: ListSatisfactionSurveys :many
SELECT * FROM satisfaction_surveys ORDER BY submitted_at DESC;

-- name: ListSatisfactionSurveysByTicketID :many
SELECT * FROM satisfaction_surveys WHERE ticket_id = $1 ORDER BY submitted_at DESC;

-- name: DeleteSatisfactionSurvey :exec
DELETE FROM satisfaction_surveys WHERE id = $1;
