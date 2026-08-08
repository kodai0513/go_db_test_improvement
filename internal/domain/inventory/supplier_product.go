package inventory

import (
	"context"

	"github.com/google/uuid"
)

// SupplierProduct は仕入先が取り扱う商品を表す。
type SupplierProduct struct {
	ID            uuid.UUID
	SupplierID    uuid.UUID
	ProductID     uuid.UUID
	LeadTimeDays  int32
}

// SupplierProductRepository は supplier_products テーブルへのアクセスを抽象化する。
type SupplierProductRepository interface {
	Create(ctx context.Context, sp *SupplierProduct) error
	Get(ctx context.Context, id uuid.UUID) (*SupplierProduct, error)
	List(ctx context.Context) ([]*SupplierProduct, error)
	ListBySupplierID(ctx context.Context, supplierID uuid.UUID) ([]*SupplierProduct, error)
	ListByProductID(ctx context.Context, productID uuid.UUID) ([]*SupplierProduct, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
