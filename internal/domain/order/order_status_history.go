package order

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// OrderStatusHistory は注文ステータスの変更履歴を表す。
type OrderStatusHistory struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	Status    string
	ChangedAt time.Time
}

// OrderStatusHistoryRepository は order_status_histories テーブルへのアクセスを抽象化する。
type OrderStatusHistoryRepository interface {
	Create(ctx context.Context, h *OrderStatusHistory) error
	Get(ctx context.Context, id uuid.UUID) (*OrderStatusHistory, error)
	List(ctx context.Context) ([]*OrderStatusHistory, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*OrderStatusHistory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
