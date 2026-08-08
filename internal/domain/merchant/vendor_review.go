package merchant

import (
	"context"

	"github.com/google/uuid"
)

// VendorReview はベンダーに対する会員のレビューを表す。
type VendorReview struct {
	ID       uuid.UUID
	VendorID uuid.UUID
	MemberID uuid.UUID
	Rating   int32
	Comment  string
}

// VendorReviewRepository は vendor_reviews テーブルへのアクセスを抽象化する。
type VendorReviewRepository interface {
	Create(ctx context.Context, r *VendorReview) error
	Get(ctx context.Context, id uuid.UUID) (*VendorReview, error)
	List(ctx context.Context) ([]*VendorReview, error)
	ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*VendorReview, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
