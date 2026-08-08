package notification

import (
	"context"

	"github.com/google/uuid"
)

// NotificationChannel は通知チャネル(メール、プッシュ等)を表す。
type NotificationChannel struct {
	ID   uuid.UUID
	Name string
}

// NotificationChannelRepository は notification_channels テーブルへのアクセスを抽象化する。
type NotificationChannelRepository interface {
	Create(ctx context.Context, c *NotificationChannel) error
	Get(ctx context.Context, id uuid.UUID) (*NotificationChannel, error)
	List(ctx context.Context) ([]*NotificationChannel, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
