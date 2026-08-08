package inventory

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// StockMovement は在庫の増減履歴を表す。
type StockMovement struct {
	ID             uuid.UUID
	InventoryID    uuid.UUID
	QuantityDelta  int32
	Reason         string
	CreatedAt      time.Time
}

// StockMovementRepository は stock_movements テーブルへのアクセスを抽象化する。
type StockMovementRepository interface {
	Create(ctx context.Context, m *StockMovement) error
	Get(ctx context.Context, id uuid.UUID) (*StockMovement, error)
	List(ctx context.Context) ([]*StockMovement, error)
	ListByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]*StockMovement, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
