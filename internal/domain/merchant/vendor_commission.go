package merchant

import (
	"context"

	"github.com/google/uuid"
)

// VendorCommission はベンダーの手数料率を表す。
type VendorCommission struct {
	ID          uuid.UUID
	VendorID    uuid.UUID
	RatePercent int32
}

// VendorCommissionRepository は vendor_commissions テーブルへのアクセスを抽象化する。
type VendorCommissionRepository interface {
	Create(ctx context.Context, c *VendorCommission) error
	Get(ctx context.Context, id uuid.UUID) (*VendorCommission, error)
	List(ctx context.Context) ([]*VendorCommission, error)
	ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*VendorCommission, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
