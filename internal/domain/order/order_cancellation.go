package order

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// OrderCancellation は注文のキャンセル情報を表す。
type OrderCancellation struct {
	ID          uuid.UUID
	OrderID     uuid.UUID
	Reason      string
	CancelledAt time.Time
}

// OrderCancellationRepository は order_cancellations テーブルへのアクセスを抽象化する。
type OrderCancellationRepository interface {
	Create(ctx context.Context, c *OrderCancellation) error
	Get(ctx context.Context, id uuid.UUID) (*OrderCancellation, error)
	List(ctx context.Context) ([]*OrderCancellation, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*OrderCancellation, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
