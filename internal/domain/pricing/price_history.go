package pricing

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// PriceHistory は商品の価格変更履歴を表す。
type PriceHistory struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	PriceYen  int32
	ChangedAt time.Time
}

// PriceHistoryRepository は price_histories テーブルへのアクセスを抽象化する。
type PriceHistoryRepository interface {
	Create(ctx context.Context, p *PriceHistory) error
	Get(ctx context.Context, id uuid.UUID) (*PriceHistory, error)
	List(ctx context.Context) ([]*PriceHistory, error)
	ListByProductID(ctx context.Context, productID uuid.UUID) ([]*PriceHistory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
