package inventory

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// InventoryAudit は倉庫の棚卸結果を表す。
type InventoryAudit struct {
	ID          uuid.UUID
	WarehouseID uuid.UUID
	PerformedAt time.Time
	Result      string
}

// InventoryAuditRepository は inventory_audits テーブルへのアクセスを抽象化する。
type InventoryAuditRepository interface {
	Create(ctx context.Context, a *InventoryAudit) error
	Get(ctx context.Context, id uuid.UUID) (*InventoryAudit, error)
	List(ctx context.Context) ([]*InventoryAudit, error)
	ListByWarehouseID(ctx context.Context, warehouseID uuid.UUID) ([]*InventoryAudit, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
