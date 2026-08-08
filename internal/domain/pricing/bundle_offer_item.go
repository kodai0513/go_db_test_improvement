package pricing

import (
	"context"

	"github.com/google/uuid"
)

// BundleOfferItem はセット販売オファーに含まれる商品を表す。
type BundleOfferItem struct {
	ID            uuid.UUID
	BundleOfferID uuid.UUID
	ProductID     uuid.UUID
	Quantity      int32
}

// BundleOfferItemRepository は bundle_offer_items テーブルへのアクセスを抽象化する。
type BundleOfferItemRepository interface {
	Create(ctx context.Context, b *BundleOfferItem) error
	Get(ctx context.Context, id uuid.UUID) (*BundleOfferItem, error)
	List(ctx context.Context) ([]*BundleOfferItem, error)
	ListByBundleOfferID(ctx context.Context, bundleOfferID uuid.UUID) ([]*BundleOfferItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
