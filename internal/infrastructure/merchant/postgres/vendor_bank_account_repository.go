package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres/sqlc"
)

// VendorBankAccountRepository は merchant.VendorBankAccountRepository の実装。
type VendorBankAccountRepository struct {
	q *sqlc.Queries
}

func NewVendorBankAccountRepository(db *sql.DB) *VendorBankAccountRepository {
	return &VendorBankAccountRepository{q: sqlc.New(db)}
}

var _ merchant.VendorBankAccountRepository = (*VendorBankAccountRepository)(nil)

func (r *VendorBankAccountRepository) Create(ctx context.Context, a *merchant.VendorBankAccount) error {
	row, err := r.q.CreateVendorBankAccount(ctx, sqlc.CreateVendorBankAccountParams{
		ID:            a.ID,
		VendorID:      a.VendorID,
		AccountNumber: a.AccountNumber,
	})
	if err != nil {
		return err
	}
	*a = toVendorBankAccountDomain(row)
	return nil
}

func (r *VendorBankAccountRepository) Get(ctx context.Context, id uuid.UUID) (*merchant.VendorBankAccount, error) {
	row, err := r.q.GetVendorBankAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toVendorBankAccountDomain(row)
	return &a, nil
}

func (r *VendorBankAccountRepository) List(ctx context.Context) ([]*merchant.VendorBankAccount, error) {
	rows, err := r.q.ListVendorBankAccounts(ctx)
	if err != nil {
		return nil, err
	}
	accounts := make([]*merchant.VendorBankAccount, 0, len(rows))
	for _, row := range rows {
		a := toVendorBankAccountDomain(row)
		accounts = append(accounts, &a)
	}
	return accounts, nil
}

func (r *VendorBankAccountRepository) ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*merchant.VendorBankAccount, error) {
	rows, err := r.q.ListVendorBankAccountsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	accounts := make([]*merchant.VendorBankAccount, 0, len(rows))
	for _, row := range rows {
		a := toVendorBankAccountDomain(row)
		accounts = append(accounts, &a)
	}
	return accounts, nil
}

func (r *VendorBankAccountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVendorBankAccount(ctx, id)
}

func toVendorBankAccountDomain(row sqlc.VendorBankAccount) merchant.VendorBankAccount {
	return merchant.VendorBankAccount{
		ID:            row.ID,
		VendorID:      row.VendorID,
		AccountNumber: row.AccountNumber,
	}
}
