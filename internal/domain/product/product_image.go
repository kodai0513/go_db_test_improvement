package product

import (
	"context"

	"github.com/google/uuid"
)

// ProductImage は商品画像を表す。
type ProductImage struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	URL       string
	SortOrder int32
}

// ProductImageRepository は product_images テーブルへのアクセスを抽象化する。
type ProductImageRepository interface {
	Create(ctx context.Context, img *ProductImage) error
	Get(ctx context.Context, id uuid.UUID) (*ProductImage, error)
	List(ctx context.Context) ([]*ProductImage, error)
	ListByProductID(ctx context.Context, productID uuid.UUID) ([]*ProductImage, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
