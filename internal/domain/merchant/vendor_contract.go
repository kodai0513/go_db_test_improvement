package merchant

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// VendorContract はベンダーとの契約期間を表す。EndsAt は未定の場合ゼロ値を許容する。
type VendorContract struct {
	ID       uuid.UUID
	VendorID uuid.UUID
	StartsAt time.Time
	EndsAt   time.Time
}

// VendorContractRepository は vendor_contracts テーブルへのアクセスを抽象化する。
type VendorContractRepository interface {
	Create(ctx context.Context, c *VendorContract) error
	Get(ctx context.Context, id uuid.UUID) (*VendorContract, error)
	List(ctx context.Context) ([]*VendorContract, error)
	ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*VendorContract, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
