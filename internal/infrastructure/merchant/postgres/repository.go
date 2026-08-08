package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres/sqlc"
)

// Repository は merchant.VendorRepository の実装。
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ merchant.VendorRepository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, v *merchant.Vendor) error {
	row, err := r.q.CreateVendor(ctx, sqlc.CreateVendorParams{
		ID:        v.ID,
		Name:      v.Name,
		Email:     v.Email,
		CreatedAt: v.CreatedAt,
	})
	if err != nil {
		return err
	}
	*v = toVendorDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*merchant.Vendor, error) {
	row, err := r.q.GetVendor(ctx, id)
	if err != nil {
		return nil, err
	}
	v := toVendorDomain(row)
	return &v, nil
}

func (r *Repository) List(ctx context.Context) ([]*merchant.Vendor, error) {
	rows, err := r.q.ListVendors(ctx)
	if err != nil {
		return nil, err
	}
	vendors := make([]*merchant.Vendor, 0, len(rows))
	for _, row := range rows {
		v := toVendorDomain(row)
		vendors = append(vendors, &v)
	}
	return vendors, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVendor(ctx, id)
}

func toVendorDomain(row sqlc.Vendor) merchant.Vendor {
	return merchant.Vendor{
		ID:        row.ID,
		Name:      row.Name,
		Email:     row.Email,
		CreatedAt: row.CreatedAt,
	}
}
