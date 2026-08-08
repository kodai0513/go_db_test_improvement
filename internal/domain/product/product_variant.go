package product

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ProductVariant は商品のバリエーション(SKU)を表す。
type ProductVariant struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	SKU       string
	PriceYen  int32
	CreatedAt time.Time
}

// ProductVariantRepository は product_variants テーブルへのアクセスを抽象化する。
type ProductVariantRepository interface {
	Create(ctx context.Context, v *ProductVariant) error
	Get(ctx context.Context, id uuid.UUID) (*ProductVariant, error)
	List(ctx context.Context) ([]*ProductVariant, error)
	ListByProductID(ctx context.Context, productID uuid.UUID) ([]*ProductVariant, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
