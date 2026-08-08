package merchant

import (
	"context"

	"github.com/google/uuid"
)

// VendorDocument はベンダーの提出書類を表す。
type VendorDocument struct {
	ID       uuid.UUID
	VendorID uuid.UUID
	URL      string
	DocType  string
}

// VendorDocumentRepository は vendor_documents テーブルへのアクセスを抽象化する。
type VendorDocumentRepository interface {
	Create(ctx context.Context, d *VendorDocument) error
	Get(ctx context.Context, id uuid.UUID) (*VendorDocument, error)
	List(ctx context.Context) ([]*VendorDocument, error)
	ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*VendorDocument, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
