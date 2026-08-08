package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres/sqlc"
)

// VendorPayoutRepository は merchant.VendorPayoutRepository の実装。
type VendorPayoutRepository struct {
	q *sqlc.Queries
}

func NewVendorPayoutRepository(db *sql.DB) *VendorPayoutRepository {
	return &VendorPayoutRepository{q: sqlc.New(db)}
}

var _ merchant.VendorPayoutRepository = (*VendorPayoutRepository)(nil)

func (r *VendorPayoutRepository) Create(ctx context.Context, p *merchant.VendorPayout) error {
	row, err := r.q.CreateVendorPayout(ctx, sqlc.CreateVendorPayoutParams{
		ID:        p.ID,
		VendorID:  p.VendorID,
		AmountYen: p.AmountYen,
		PaidAt:    p.PaidAt,
	})
	if err != nil {
		return err
	}
	*p = toVendorPayoutDomain(row)
	return nil
}

func (r *VendorPayoutRepository) Get(ctx context.Context, id uuid.UUID) (*merchant.VendorPayout, error) {
	row, err := r.q.GetVendorPayout(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toVendorPayoutDomain(row)
	return &p, nil
}

func (r *VendorPayoutRepository) List(ctx context.Context) ([]*merchant.VendorPayout, error) {
	rows, err := r.q.ListVendorPayouts(ctx)
	if err != nil {
		return nil, err
	}
	payouts := make([]*merchant.VendorPayout, 0, len(rows))
	for _, row := range rows {
		p := toVendorPayoutDomain(row)
		payouts = append(payouts, &p)
	}
	return payouts, nil
}

func (r *VendorPayoutRepository) ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*merchant.VendorPayout, error) {
	rows, err := r.q.ListVendorPayoutsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	payouts := make([]*merchant.VendorPayout, 0, len(rows))
	for _, row := range rows {
		p := toVendorPayoutDomain(row)
		payouts = append(payouts, &p)
	}
	return payouts, nil
}

func (r *VendorPayoutRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVendorPayout(ctx, id)
}

func toVendorPayoutDomain(row sqlc.VendorPayout) merchant.VendorPayout {
	return merchant.VendorPayout{
		ID:        row.ID,
		VendorID:  row.VendorID,
		AmountYen: row.AmountYen,
		PaidAt:    row.PaidAt,
	}
}
