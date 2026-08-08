package notification

import (
	"context"

	"github.com/google/uuid"
)

// NotificationTemplate は通知テンプレートを表す。
type NotificationTemplate struct {
	ID   uuid.UUID
	Name string
	Body string
}

// NotificationTemplateRepository は notification_templates テーブルへのアクセスを抽象化する。
type NotificationTemplateRepository interface {
	Create(ctx context.Context, t *NotificationTemplate) error
	Get(ctx context.Context, id uuid.UUID) (*NotificationTemplate, error)
	List(ctx context.Context) ([]*NotificationTemplate, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
