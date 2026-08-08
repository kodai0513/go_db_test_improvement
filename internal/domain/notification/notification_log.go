package notification

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// NotificationLog は通知の送信結果ログを表す。
type NotificationLog struct {
	ID             uuid.UUID
	NotificationID uuid.UUID
	Status         string
	SentAt         time.Time
}

// NotificationLogRepository は notification_logs テーブルへのアクセスを抽象化する。
type NotificationLogRepository interface {
	Create(ctx context.Context, l *NotificationLog) error
	Get(ctx context.Context, id uuid.UUID) (*NotificationLog, error)
	List(ctx context.Context) ([]*NotificationLog, error)
	ListByNotificationID(ctx context.Context, notificationID uuid.UUID) ([]*NotificationLog, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
