package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Invoice は注文に対して発行される請求書を表す。
type Invoice struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	AmountYen int32
	IssuedAt  time.Time
}

// InvoiceRepository は invoices テーブルへのアクセスを抽象化する。
type InvoiceRepository interface {
	Create(ctx context.Context, i *Invoice) error
	Get(ctx context.Context, id uuid.UUID) (*Invoice, error)
	List(ctx context.Context) ([]*Invoice, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*Invoice, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
