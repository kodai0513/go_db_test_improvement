package support

import (
	"context"

	"github.com/google/uuid"
)

// TicketCategory はチケットの分類カテゴリを表す。
type TicketCategory struct {
	ID   uuid.UUID
	Name string
}

// TicketCategoryRepository は ticket_categories テーブルへのアクセスを抽象化する。
type TicketCategoryRepository interface {
	Create(ctx context.Context, c *TicketCategory) error
	Get(ctx context.Context, id uuid.UUID) (*TicketCategory, error)
	List(ctx context.Context) ([]*TicketCategory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
