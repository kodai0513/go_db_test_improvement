package inventory

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Warehouse は倉庫を表す。
type Warehouse struct {
	ID        uuid.UUID
	Name      string
	Address   string
	CreatedAt time.Time
}

// WarehouseRepository は warehouses テーブルへのアクセスを抽象化する。
type WarehouseRepository interface {
	Create(ctx context.Context, w *Warehouse) error
	Get(ctx context.Context, id uuid.UUID) (*Warehouse, error)
	List(ctx context.Context) ([]*Warehouse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
