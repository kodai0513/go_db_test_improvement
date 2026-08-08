package pricing

import (
	"context"

	"github.com/google/uuid"
)

// PriceRule は価格表に紐づく商品ごとの価格ルールを表す。
type PriceRule struct {
	ID          uuid.UUID
	PriceListID uuid.UUID
	ProductID   uuid.UUID
	PriceYen    int32
}

// PriceRuleRepository は price_rules テーブルへのアクセスを抽象化する。
type PriceRuleRepository interface {
	Create(ctx context.Context, p *PriceRule) error
	Get(ctx context.Context, id uuid.UUID) (*PriceRule, error)
	List(ctx context.Context) ([]*PriceRule, error)
	ListByPriceListID(ctx context.Context, priceListID uuid.UUID) ([]*PriceRule, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
