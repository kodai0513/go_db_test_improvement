package inventory

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// PurchaseOrder は仕入先への発注を表す。
type PurchaseOrder struct {
	ID         uuid.UUID
	SupplierID uuid.UUID
	Status     string
	CreatedAt  time.Time
}

// PurchaseOrderRepository は purchase_orders テーブルへのアクセスを抽象化する。
type PurchaseOrderRepository interface {
	Create(ctx context.Context, p *PurchaseOrder) error
	Get(ctx context.Context, id uuid.UUID) (*PurchaseOrder, error)
	List(ctx context.Context) ([]*PurchaseOrder, error)
	ListBySupplierID(ctx context.Context, supplierID uuid.UUID) ([]*PurchaseOrder, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
