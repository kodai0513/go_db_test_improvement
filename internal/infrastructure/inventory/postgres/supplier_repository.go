package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres/sqlc"
)

// SupplierRepository は inventory.SupplierRepository の実装。
type SupplierRepository struct {
	q *sqlc.Queries
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{q: sqlc.New(db)}
}

var _ inventory.SupplierRepository = (*SupplierRepository)(nil)

func (r *SupplierRepository) Create(ctx context.Context, s *inventory.Supplier) error {
	row, err := r.q.CreateSupplier(ctx, sqlc.CreateSupplierParams{
		ID:           s.ID,
		Name:         s.Name,
		ContactEmail: s.ContactEmail,
		CreatedAt:    s.CreatedAt,
	})
	if err != nil {
		return err
	}
	*s = toSupplierDomain(row)
	return nil
}

func (r *SupplierRepository) Get(ctx context.Context, id uuid.UUID) (*inventory.Supplier, error) {
	row, err := r.q.GetSupplier(ctx, id)
	if err != nil {
		return nil, err
	}
	s := toSupplierDomain(row)
	return &s, nil
}

func (r *SupplierRepository) List(ctx context.Context) ([]*inventory.Supplier, error) {
	rows, err := r.q.ListSuppliers(ctx)
	if err != nil {
		return nil, err
	}
	suppliers := make([]*inventory.Supplier, 0, len(rows))
	for _, row := range rows {
		s := toSupplierDomain(row)
		suppliers = append(suppliers, &s)
	}
	return suppliers, nil
}

func (r *SupplierRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteSupplier(ctx, id)
}

func toSupplierDomain(row sqlc.Supplier) inventory.Supplier {
	return inventory.Supplier{
		ID:           row.ID,
		Name:         row.Name,
		ContactEmail: row.ContactEmail,
		CreatedAt:    row.CreatedAt,
	}
}
