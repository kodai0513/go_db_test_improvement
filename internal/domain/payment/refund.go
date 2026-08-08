package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Refund は決済に対する返金を表す。
type Refund struct {
	ID         uuid.UUID
	PaymentID  uuid.UUID
	AmountYen  int32
	Reason     string
	RefundedAt time.Time
}

// RefundRepository は refunds テーブルへのアクセスを抽象化する。
type RefundRepository interface {
	Create(ctx context.Context, r *Refund) error
	Get(ctx context.Context, id uuid.UUID) (*Refund, error)
	List(ctx context.Context) ([]*Refund, error)
	ListByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*Refund, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
