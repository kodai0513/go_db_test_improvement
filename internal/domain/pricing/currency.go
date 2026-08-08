package pricing

import (
	"context"

	"github.com/google/uuid"
)

// Currency は通貨を表す。
type Currency struct {
	ID   uuid.UUID
	Code string
}

// CurrencyRepository は currencies テーブルへのアクセスを抽象化する。
type CurrencyRepository interface {
	Create(ctx context.Context, c *Currency) error
	Get(ctx context.Context, id uuid.UUID) (*Currency, error)
	List(ctx context.Context) ([]*Currency, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
