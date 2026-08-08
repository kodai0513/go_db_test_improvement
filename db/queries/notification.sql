-- name: CreateNotification :one
INSERT INTO notifications (id, member_id, message, is_read, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetNotification :one
SELECT * FROM notifications WHERE id = $1;

-- name: ListNotifications :many
SELECT * FROM notifications ORDER BY created_at DESC;

-- name: CreateNotificationTemplate :one
INSERT INTO notification_templates (id, name, body)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetNotificationTemplate :one
SELECT * FROM notification_templates WHERE id = $1;

-- name: ListNotificationTemplates :many
SELECT * FROM notification_templates ORDER BY id DESC;

-- name: DeleteNotificationTemplate :exec
DELETE FROM notification_templates WHERE id = $1;

-- name: CreateNotificationChannel :one
INSERT INTO notification_channels (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetNotificationChannel :one
SELECT * FROM notification_channels WHERE id = $1;

-- name: ListNotificationChannels :many
SELECT * FROM notification_channels ORDER BY id DESC;

-- name: DeleteNotificationChannel :exec
DELETE FROM notification_channels WHERE id = $1;

-- name: CreateNotificationPreference :one
INSERT INTO notification_preferences (id, member_id, channel_id, enabled)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetNotificationPreference :one
SELECT * FROM notification_preferences WHERE id = $1;

-- name: ListNotificationPreferences :many
SELECT * FROM notification_preferences ORDER BY id DESC;

-- name: ListNotificationPreferencesByMemberID :many
SELECT * FROM notification_preferences WHERE member_id = $1 ORDER BY id DESC;

-- name: ListNotificationPreferencesByChannelID :many
SELECT * FROM notification_preferences WHERE channel_id = $1 ORDER BY id DESC;

-- name: DeleteNotificationPreference :exec
DELETE FROM notification_preferences WHERE id = $1;

-- name: CreateNotificationLog :one
INSERT INTO notification_logs (id, notification_id, status, sent_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetNotificationLog :one
SELECT * FROM notification_logs WHERE id = $1;

-- name: ListNotificationLogs :many
SELECT * FROM notification_logs ORDER BY sent_at DESC;

-- name: ListNotificationLogsByNotificationID :many
SELECT * FROM notification_logs WHERE notification_id = $1 ORDER BY sent_at DESC;

-- name: DeleteNotificationLog :exec
DELETE FROM notification_logs WHERE id = $1;

-- name: CreateNotificationSubscription :one
INSERT INTO notification_subscriptions (id, member_id, topic, subscribed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetNotificationSubscription :one
SELECT * FROM notification_subscriptions WHERE id = $1;

-- name: ListNotificationSubscriptions :many
SELECT * FROM notification_subscriptions ORDER BY subscribed_at DESC;

-- name: ListNotificationSubscriptionsByMemberID :many
SELECT * FROM notification_subscriptions WHERE member_id = $1 ORDER BY subscribed_at DESC;

-- name: DeleteNotificationSubscription :exec
DELETE FROM notification_subscriptions WHERE id = $1;

-- name: CreatePushDevice :one
INSERT INTO push_devices (id, member_id, device_token, registered_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPushDevice :one
SELECT * FROM push_devices WHERE id = $1;

-- name: ListPushDevices :many
SELECT * FROM push_devices ORDER BY registered_at DESC;

-- name: ListPushDevicesByMemberID :many
SELECT * FROM push_devices WHERE member_id = $1 ORDER BY registered_at DESC;

-- name: DeletePushDevice :exec
DELETE FROM push_devices WHERE id = $1;

-- name: CreateEmailLog :one
INSERT INTO email_logs (id, member_id, subject, sent_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetEmailLog :one
SELECT * FROM email_logs WHERE id = $1;

-- name: ListEmailLogs :many
SELECT * FROM email_logs ORDER BY sent_at DESC;

-- name: ListEmailLogsByMemberID :many
SELECT * FROM email_logs WHERE member_id = $1 ORDER BY sent_at DESC;

-- name: DeleteEmailLog :exec
DELETE FROM email_logs WHERE id = $1;

-- name: CreateSMSLog :one
INSERT INTO sms_logs (id, member_id, body, sent_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetSMSLog :one
SELECT * FROM sms_logs WHERE id = $1;

-- name: ListSMSLogs :many
SELECT * FROM sms_logs ORDER BY sent_at DESC;

-- name: ListSMSLogsByMemberID :many
SELECT * FROM sms_logs WHERE member_id = $1 ORDER BY sent_at DESC;

-- name: DeleteSMSLog :exec
DELETE FROM sms_logs WHERE id = $1;

-- name: CreateNotificationBatch :one
INSERT INTO notification_batches (id, template_id, status, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetNotificationBatch :one
SELECT * FROM notification_batches WHERE id = $1;

-- name: ListNotificationBatches :many
SELECT * FROM notification_batches ORDER BY created_at DESC;

-- name: ListNotificationBatchesByTemplateID :many
SELECT * FROM notification_batches WHERE template_id = $1 ORDER BY created_at DESC;

-- name: DeleteNotificationBatch :exec
DELETE FROM notification_batches WHERE id = $1;
