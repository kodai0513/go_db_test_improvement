package inventory

import (
	"context"

	"github.com/google/uuid"
)

// PurchaseOrderItem は発注明細を表す。
type PurchaseOrderItem struct {
	ID              uuid.UUID
	PurchaseOrderID uuid.UUID
	ProductID       uuid.UUID
	Quantity        int32
	UnitPriceYen    int32
}

// PurchaseOrderItemRepository は purchase_order_items テーブルへのアクセスを抽象化する。
type PurchaseOrderItemRepository interface {
	Create(ctx context.Context, i *PurchaseOrderItem) error
	Get(ctx context.Context, id uuid.UUID) (*PurchaseOrderItem, error)
	List(ctx context.Context) ([]*PurchaseOrderItem, error)
	ListByPurchaseOrderID(ctx context.Context, purchaseOrderID uuid.UUID) ([]*PurchaseOrderItem, error)
	ListByProductID(ctx context.Context, productID uuid.UUID) ([]*PurchaseOrderItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
