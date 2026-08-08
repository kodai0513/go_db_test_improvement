package product

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ProductCategory は商品とカテゴリの中間テーブルを表す。
type ProductCategory struct {
	ID         uuid.UUID
	ProductID  uuid.UUID
	CategoryID uuid.UUID
	CreatedAt  time.Time
}

// ProductCategoryRepository は product_categories テーブルへのアクセスを抽象化する。
type ProductCategoryRepository interface {
	Create(ctx context.Context, pc *ProductCategory) error
	Get(ctx context.Context, id uuid.UUID) (*ProductCategory, error)
	List(ctx context.Context) ([]*ProductCategory, error)
	ListByProductID(ctx context.Context, productID uuid.UUID) ([]*ProductCategory, error)
	ListByCategoryID(ctx context.Context, categoryID uuid.UUID) ([]*ProductCategory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
