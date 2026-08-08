package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// PaymentMethod は会員が登録した支払い方法を表す。
type PaymentMethod struct {
	ID        uuid.UUID
	MemberID  uuid.UUID
	Type      string
	IsDefault bool
	CreatedAt time.Time
}

// PaymentMethodRepository は payment_methods テーブルへのアクセスを抽象化する。
type PaymentMethodRepository interface {
	Create(ctx context.Context, p *PaymentMethod) error
	Get(ctx context.Context, id uuid.UUID) (*PaymentMethod, error)
	List(ctx context.Context) ([]*PaymentMethod, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*PaymentMethod, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
