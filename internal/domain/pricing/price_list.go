package pricing

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// PriceList は価格表を表す。
type PriceList struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

// PriceListRepository は price_lists テーブルへのアクセスを抽象化する。
type PriceListRepository interface {
	Create(ctx context.Context, p *PriceList) error
	Get(ctx context.Context, id uuid.UUID) (*PriceList, error)
	List(ctx context.Context) ([]*PriceList, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
