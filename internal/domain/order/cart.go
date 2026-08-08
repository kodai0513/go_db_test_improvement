package order

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Cart は会員が保持するショッピングカートを表す。
type Cart struct {
	ID        uuid.UUID
	MemberID  uuid.UUID
	CreatedAt time.Time
}

// CartRepository は carts テーブルへのアクセスを抽象化する。
type CartRepository interface {
	Create(ctx context.Context, c *Cart) error
	Get(ctx context.Context, id uuid.UUID) (*Cart, error)
	List(ctx context.Context) ([]*Cart, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*Cart, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
