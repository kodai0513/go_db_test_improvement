package content

import (
	"context"

	"github.com/google/uuid"
)

// Page は静的コンテンツページを表す。
type Page struct {
	ID    uuid.UUID
	Slug  string
	Title string
	Body  string
}

// PageRepository は pages テーブルへのアクセスを抽象化する。
type PageRepository interface {
	Create(ctx context.Context, p *Page) error
	Get(ctx context.Context, id uuid.UUID) (*Page, error)
	List(ctx context.Context) ([]*Page, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
