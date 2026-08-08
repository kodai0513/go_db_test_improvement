package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres/sqlc"
)

// VendorCommissionRepository は merchant.VendorCommissionRepository の実装。
type VendorCommissionRepository struct {
	q *sqlc.Queries
}

func NewVendorCommissionRepository(db *sql.DB) *VendorCommissionRepository {
	return &VendorCommissionRepository{q: sqlc.New(db)}
}

var _ merchant.VendorCommissionRepository = (*VendorCommissionRepository)(nil)

func (r *VendorCommissionRepository) Create(ctx context.Context, c *merchant.VendorCommission) error {
	row, err := r.q.CreateVendorCommission(ctx, sqlc.CreateVendorCommissionParams{
		ID:          c.ID,
		VendorID:    c.VendorID,
		RatePercent: c.RatePercent,
	})
	if err != nil {
		return err
	}
	*c = toVendorCommissionDomain(row)
	return nil
}

func (r *VendorCommissionRepository) Get(ctx context.Context, id uuid.UUID) (*merchant.VendorCommission, error) {
	row, err := r.q.GetVendorCommission(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toVendorCommissionDomain(row)
	return &c, nil
}

func (r *VendorCommissionRepository) List(ctx context.Context) ([]*merchant.VendorCommission, error) {
	rows, err := r.q.ListVendorCommissions(ctx)
	if err != nil {
		return nil, err
	}
	commissions := make([]*merchant.VendorCommission, 0, len(rows))
	for _, row := range rows {
		c := toVendorCommissionDomain(row)
		commissions = append(commissions, &c)
	}
	return commissions, nil
}

func (r *VendorCommissionRepository) ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*merchant.VendorCommission, error) {
	rows, err := r.q.ListVendorCommissionsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	commissions := make([]*merchant.VendorCommission, 0, len(rows))
	for _, row := range rows {
		c := toVendorCommissionDomain(row)
		commissions = append(commissions, &c)
	}
	return commissions, nil
}

func (r *VendorCommissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVendorCommission(ctx, id)
}

func toVendorCommissionDomain(row sqlc.VendorCommission) merchant.VendorCommission {
	return merchant.VendorCommission{
		ID:          row.ID,
		VendorID:    row.VendorID,
		RatePercent: row.RatePercent,
	}
}
