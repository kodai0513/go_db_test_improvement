package merchant

import (
	"context"

	"github.com/google/uuid"
)

// VendorProduct はベンダーが取り扱う商品を表す。
type VendorProduct struct {
	ID        uuid.UUID
	VendorID  uuid.UUID
	ProductID uuid.UUID
}

// VendorProductRepository は vendor_products テーブルへのアクセスを抽象化する。
type VendorProductRepository interface {
	Create(ctx context.Context, p *VendorProduct) error
	Get(ctx context.Context, id uuid.UUID) (*VendorProduct, error)
	List(ctx context.Context) ([]*VendorProduct, error)
	ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*VendorProduct, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
