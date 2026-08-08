package merchant

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// VendorPayout はベンダーへの支払いを表す。
type VendorPayout struct {
	ID        uuid.UUID
	VendorID  uuid.UUID
	AmountYen int32
	PaidAt    time.Time
}

// VendorPayoutRepository は vendor_payouts テーブルへのアクセスを抽象化する。
type VendorPayoutRepository interface {
	Create(ctx context.Context, p *VendorPayout) error
	Get(ctx context.Context, id uuid.UUID) (*VendorPayout, error)
	List(ctx context.Context) ([]*VendorPayout, error)
	ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*VendorPayout, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
