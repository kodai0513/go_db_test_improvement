package product

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Category は商品カテゴリを表す。
type Category struct {
	ID               uuid.UUID
	Name             string
	ParentCategoryID uuid.UUID
	CreatedAt        time.Time
}

// CategoryRepository は categories テーブルへのアクセスを抽象化する。
type CategoryRepository interface {
	Create(ctx context.Context, c *Category) error
	Get(ctx context.Context, id uuid.UUID) (*Category, error)
	List(ctx context.Context) ([]*Category, error)
	ListByParentCategoryID(ctx context.Context, parentCategoryID uuid.UUID) ([]*Category, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
