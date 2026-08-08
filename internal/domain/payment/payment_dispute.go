package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// PaymentDispute は決済に対する異議申し立て(チャージバック等)を表す。
type PaymentDispute struct {
	ID        uuid.UUID
	PaymentID uuid.UUID
	Reason    string
	Status    string
	CreatedAt time.Time
}

// PaymentDisputeRepository は payment_disputes テーブルへのアクセスを抽象化する。
type PaymentDisputeRepository interface {
	Create(ctx context.Context, d *PaymentDispute) error
	Get(ctx context.Context, id uuid.UUID) (*PaymentDispute, error)
	List(ctx context.Context) ([]*PaymentDispute, error)
	ListByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*PaymentDispute, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
