package order

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// OrderNote は注文に付与される備考メモを表す。
type OrderNote struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	Note      string
	CreatedAt time.Time
}

// OrderNoteRepository は order_notes テーブルへのアクセスを抽象化する。
type OrderNoteRepository interface {
	Create(ctx context.Context, n *OrderNote) error
	Get(ctx context.Context, id uuid.UUID) (*OrderNote, error)
	List(ctx context.Context) ([]*OrderNote, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*OrderNote, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
