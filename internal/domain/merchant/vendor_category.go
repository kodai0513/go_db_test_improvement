package merchant

import (
	"context"

	"github.com/google/uuid"
)

// VendorCategory はベンダーが所属するカテゴリを表す。
type VendorCategory struct {
	ID         uuid.UUID
	VendorID   uuid.UUID
	CategoryID uuid.UUID
}

// VendorCategoryRepository は vendor_categories テーブルへのアクセスを抽象化する。
type VendorCategoryRepository interface {
	Create(ctx context.Context, c *VendorCategory) error
	Get(ctx context.Context, id uuid.UUID) (*VendorCategory, error)
	List(ctx context.Context) ([]*VendorCategory, error)
	ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*VendorCategory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
