package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres/sqlc"
)

// VendorDocumentRepository は merchant.VendorDocumentRepository の実装。
type VendorDocumentRepository struct {
	q *sqlc.Queries
}

func NewVendorDocumentRepository(db *sql.DB) *VendorDocumentRepository {
	return &VendorDocumentRepository{q: sqlc.New(db)}
}

var _ merchant.VendorDocumentRepository = (*VendorDocumentRepository)(nil)

func (r *VendorDocumentRepository) Create(ctx context.Context, d *merchant.VendorDocument) error {
	row, err := r.q.CreateVendorDocument(ctx, sqlc.CreateVendorDocumentParams{
		ID:       d.ID,
		VendorID: d.VendorID,
		Url:      d.URL,
		DocType:  d.DocType,
	})
	if err != nil {
		return err
	}
	*d = toVendorDocumentDomain(row)
	return nil
}

func (r *VendorDocumentRepository) Get(ctx context.Context, id uuid.UUID) (*merchant.VendorDocument, error) {
	row, err := r.q.GetVendorDocument(ctx, id)
	if err != nil {
		return nil, err
	}
	d := toVendorDocumentDomain(row)
	return &d, nil
}

func (r *VendorDocumentRepository) List(ctx context.Context) ([]*merchant.VendorDocument, error) {
	rows, err := r.q.ListVendorDocuments(ctx)
	if err != nil {
		return nil, err
	}
	documents := make([]*merchant.VendorDocument, 0, len(rows))
	for _, row := range rows {
		d := toVendorDocumentDomain(row)
		documents = append(documents, &d)
	}
	return documents, nil
}

func (r *VendorDocumentRepository) ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*merchant.VendorDocument, error) {
	rows, err := r.q.ListVendorDocumentsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	documents := make([]*merchant.VendorDocument, 0, len(rows))
	for _, row := range rows {
		d := toVendorDocumentDomain(row)
		documents = append(documents, &d)
	}
	return documents, nil
}

func (r *VendorDocumentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVendorDocument(ctx, id)
}

func toVendorDocumentDomain(row sqlc.VendorDocument) merchant.VendorDocument {
	return merchant.VendorDocument{
		ID:       row.ID,
		VendorID: row.VendorID,
		URL:      row.Url,
		DocType:  row.DocType,
	}
}
