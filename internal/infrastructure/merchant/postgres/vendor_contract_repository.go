package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres/sqlc"
)

// VendorContractRepository は merchant.VendorContractRepository の実装。
type VendorContractRepository struct {
	q *sqlc.Queries
}

func NewVendorContractRepository(db *sql.DB) *VendorContractRepository {
	return &VendorContractRepository{q: sqlc.New(db)}
}

var _ merchant.VendorContractRepository = (*VendorContractRepository)(nil)

func (r *VendorContractRepository) Create(ctx context.Context, c *merchant.VendorContract) error {
	row, err := r.q.CreateVendorContract(ctx, sqlc.CreateVendorContractParams{
		ID:       c.ID,
		VendorID: c.VendorID,
		StartsAt: c.StartsAt,
		EndsAt:   c.EndsAt,
	})
	if err != nil {
		return err
	}
	*c = toVendorContractDomain(row)
	return nil
}

func (r *VendorContractRepository) Get(ctx context.Context, id uuid.UUID) (*merchant.VendorContract, error) {
	row, err := r.q.GetVendorContract(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toVendorContractDomain(row)
	return &c, nil
}

func (r *VendorContractRepository) List(ctx context.Context) ([]*merchant.VendorContract, error) {
	rows, err := r.q.ListVendorContracts(ctx)
	if err != nil {
		return nil, err
	}
	contracts := make([]*merchant.VendorContract, 0, len(rows))
	for _, row := range rows {
		c := toVendorContractDomain(row)
		contracts = append(contracts, &c)
	}
	return contracts, nil
}

func (r *VendorContractRepository) ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*merchant.VendorContract, error) {
	rows, err := r.q.ListVendorContractsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	contracts := make([]*merchant.VendorContract, 0, len(rows))
	for _, row := range rows {
		c := toVendorContractDomain(row)
		contracts = append(contracts, &c)
	}
	return contracts, nil
}

func (r *VendorContractRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVendorContract(ctx, id)
}

func toVendorContractDomain(row sqlc.VendorContract) merchant.VendorContract {
	return merchant.VendorContract{
		ID:       row.ID,
		VendorID: row.VendorID,
		StartsAt: row.StartsAt,
		EndsAt:   row.EndsAt,
	}
}
