package inventory

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// StockAdjustment は在庫の手動調整を表す。
// note は DB上は NULL許容だが、Go側ではゼロ値(空文字列)で許容する。
type StockAdjustment struct {
	ID               uuid.UUID
	InventoryID      uuid.UUID
	AdjustedQuantity int32
	Note             string
	CreatedAt        time.Time
}

// StockAdjustmentRepository は stock_adjustments テーブルへのアクセスを抽象化する。
type StockAdjustmentRepository interface {
	Create(ctx context.Context, a *StockAdjustment) error
	Get(ctx context.Context, id uuid.UUID) (*StockAdjustment, error)
	List(ctx context.Context) ([]*StockAdjustment, error)
	ListByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]*StockAdjustment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
