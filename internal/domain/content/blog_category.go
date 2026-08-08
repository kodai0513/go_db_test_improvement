package content

import (
	"context"

	"github.com/google/uuid"
)

// BlogCategory はブログ記事のカテゴリを表す。
type BlogCategory struct {
	ID   uuid.UUID
	Name string
}

// BlogCategoryRepository は blog_categories テーブルへのアクセスを抽象化する。
type BlogCategoryRepository interface {
	Create(ctx context.Context, c *BlogCategory) error
	Get(ctx context.Context, id uuid.UUID) (*BlogCategory, error)
	List(ctx context.Context) ([]*BlogCategory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
