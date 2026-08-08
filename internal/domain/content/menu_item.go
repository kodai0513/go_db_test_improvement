package content

import (
	"context"

	"github.com/google/uuid"
)

// MenuItem はメニューに属する個々のリンク項目を表す。
type MenuItem struct {
	ID        uuid.UUID
	MenuID    uuid.UUID
	Label     string
	URL       string
	SortOrder int32
}

// MenuItemRepository は menu_items テーブルへのアクセスを抽象化する。
type MenuItemRepository interface {
	Create(ctx context.Context, m *MenuItem) error
	Get(ctx context.Context, id uuid.UUID) (*MenuItem, error)
	List(ctx context.Context) ([]*MenuItem, error)
	ListByMenuID(ctx context.Context, menuID uuid.UUID) ([]*MenuItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
