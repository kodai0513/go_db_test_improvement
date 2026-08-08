package payment

import (
	"context"

	"github.com/google/uuid"
)

// PaymentProvider は決済代行事業者を表す。
type PaymentProvider struct {
	ID   uuid.UUID
	Name string
}

// PaymentProviderRepository は payment_providers テーブルへのアクセスを抽象化する。
type PaymentProviderRepository interface {
	Create(ctx context.Context, p *PaymentProvider) error
	Get(ctx context.Context, id uuid.UUID) (*PaymentProvider, error)
	List(ctx context.Context) ([]*PaymentProvider, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
