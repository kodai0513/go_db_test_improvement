package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres/sqlc"
)

// SupplierProductRepository は inventory.SupplierProductRepository の実装。
type SupplierProductRepository struct {
	q *sqlc.Queries
}

func NewSupplierProductRepository(db *sql.DB) *SupplierProductRepository {
	return &SupplierProductRepository{q: sqlc.New(db)}
}

var _ inventory.SupplierProductRepository = (*SupplierProductRepository)(nil)

func (r *SupplierProductRepository) Create(ctx context.Context, sp *inventory.SupplierProduct) error {
	row, err := r.q.CreateSupplierProduct(ctx, sqlc.CreateSupplierProductParams{
		ID:           sp.ID,
		SupplierID:   sp.SupplierID,
		ProductID:    sp.ProductID,
		LeadTimeDays: sp.LeadTimeDays,
	})
	if err != nil {
		return err
	}
	*sp = toSupplierProductDomain(row)
	return nil
}

func (r *SupplierProductRepository) Get(ctx context.Context, id uuid.UUID) (*inventory.SupplierProduct, error) {
	row, err := r.q.GetSupplierProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	sp := toSupplierProductDomain(row)
	return &sp, nil
}

func (r *SupplierProductRepository) List(ctx context.Context) ([]*inventory.SupplierProduct, error) {
	rows, err := r.q.ListSupplierProducts(ctx)
	if err != nil {
		return nil, err
	}
	products := make([]*inventory.SupplierProduct, 0, len(rows))
	for _, row := range rows {
		sp := toSupplierProductDomain(row)
		products = append(products, &sp)
	}
	return products, nil
}

func (r *SupplierProductRepository) ListBySupplierID(ctx context.Context, supplierID uuid.UUID) ([]*inventory.SupplierProduct, error) {
	rows, err := r.q.ListSupplierProductsBySupplierID(ctx, supplierID)
	if err != nil {
		return nil, err
	}
	products := make([]*inventory.SupplierProduct, 0, len(rows))
	for _, row := range rows {
		sp := toSupplierProductDomain(row)
		products = append(products, &sp)
	}
	return products, nil
}

func (r *SupplierProductRepository) ListByProductID(ctx context.Context, productID uuid.UUID) ([]*inventory.SupplierProduct, error) {
	rows, err := r.q.ListSupplierProductsByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	products := make([]*inventory.SupplierProduct, 0, len(rows))
	for _, row := range rows {
		sp := toSupplierProductDomain(row)
		products = append(products, &sp)
	}
	return products, nil
}

func (r *SupplierProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteSupplierProduct(ctx, id)
}

func toSupplierProductDomain(row sqlc.SupplierProduct) inventory.SupplierProduct {
	return inventory.SupplierProduct{
		ID:           row.ID,
		SupplierID:   row.SupplierID,
		ProductID:    row.ProductID,
		LeadTimeDays: row.LeadTimeDays,
	}
}
