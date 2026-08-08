package notification

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// NotificationBatch はテンプレートに基づく一括通知の実行を表す。
type NotificationBatch struct {
	ID         uuid.UUID
	TemplateID uuid.UUID
	Status     string
	CreatedAt  time.Time
}

// NotificationBatchRepository は notification_batches テーブルへのアクセスを抽象化する。
type NotificationBatchRepository interface {
	Create(ctx context.Context, b *NotificationBatch) error
	Get(ctx context.Context, id uuid.UUID) (*NotificationBatch, error)
	List(ctx context.Context) ([]*NotificationBatch, error)
	ListByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*NotificationBatch, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
