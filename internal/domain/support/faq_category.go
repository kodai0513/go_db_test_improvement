package support

import (
	"context"

	"github.com/google/uuid"
)

// FaqCategory はFAQの分類カテゴリを表す。
type FaqCategory struct {
	ID   uuid.UUID
	Name string
}

// FaqCategoryRepository は faq_categories テーブルへのアクセスを抽象化する。
type FaqCategoryRepository interface {
	Create(ctx context.Context, c *FaqCategory) error
	Get(ctx context.Context, id uuid.UUID) (*FaqCategory, error)
	List(ctx context.Context) ([]*FaqCategory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
