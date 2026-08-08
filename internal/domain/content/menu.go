package content

import (
	"context"

	"github.com/google/uuid"
)

// Menu はサイト内のナビゲーションメニューを表す。
type Menu struct {
	ID   uuid.UUID
	Name string
}

// MenuRepository は menus テーブルへのアクセスを抽象化する。
type MenuRepository interface {
	Create(ctx context.Context, m *Menu) error
	Get(ctx context.Context, id uuid.UUID) (*Menu, error)
	List(ctx context.Context) ([]*Menu, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
