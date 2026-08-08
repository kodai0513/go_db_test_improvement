package pricing

import (
	"context"

	"github.com/google/uuid"
)

// BundleOffer はセット販売オファーを表す。
type BundleOffer struct {
	ID       uuid.UUID
	Name     string
	PriceYen int32
}

// BundleOfferRepository は bundle_offers テーブルへのアクセスを抽象化する。
type BundleOfferRepository interface {
	Create(ctx context.Context, b *BundleOffer) error
	Get(ctx context.Context, id uuid.UUID) (*BundleOffer, error)
	List(ctx context.Context) ([]*BundleOffer, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
