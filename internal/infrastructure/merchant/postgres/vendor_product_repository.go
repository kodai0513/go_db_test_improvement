package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres/sqlc"
)

// VendorProductRepository は merchant.VendorProductRepository の実装。
type VendorProductRepository struct {
	q *sqlc.Queries
}

func NewVendorProductRepository(db *sql.DB) *VendorProductRepository {
	return &VendorProductRepository{q: sqlc.New(db)}
}

var _ merchant.VendorProductRepository = (*VendorProductRepository)(nil)

func (r *VendorProductRepository) Create(ctx context.Context, p *merchant.VendorProduct) error {
	row, err := r.q.CreateVendorProduct(ctx, sqlc.CreateVendorProductParams{
		ID:        p.ID,
		VendorID:  p.VendorID,
		ProductID: p.ProductID,
	})
	if err != nil {
		return err
	}
	*p = toVendorProductDomain(row)
	return nil
}

func (r *VendorProductRepository) Get(ctx context.Context, id uuid.UUID) (*merchant.VendorProduct, error) {
	row, err := r.q.GetVendorProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toVendorProductDomain(row)
	return &p, nil
}

func (r *VendorProductRepository) List(ctx context.Context) ([]*merchant.VendorProduct, error) {
	rows, err := r.q.ListVendorProducts(ctx)
	if err != nil {
		return nil, err
	}
	products := make([]*merchant.VendorProduct, 0, len(rows))
	for _, row := range rows {
		p := toVendorProductDomain(row)
		products = append(products, &p)
	}
	return products, nil
}

func (r *VendorProductRepository) ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*merchant.VendorProduct, error) {
	rows, err := r.q.ListVendorProductsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	products := make([]*merchant.VendorProduct, 0, len(rows))
	for _, row := range rows {
		p := toVendorProductDomain(row)
		products = append(products, &p)
	}
	return products, nil
}

func (r *VendorProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVendorProduct(ctx, id)
}

func toVendorProductDomain(row sqlc.VendorProduct) merchant.VendorProduct {
	return merchant.VendorProduct{
		ID:        row.ID,
		VendorID:  row.VendorID,
		ProductID: row.ProductID,
	}
}
