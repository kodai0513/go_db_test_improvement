package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres/sqlc"
)

// PaymentProviderAccountRepository は payment.PaymentProviderAccountRepository の実装。
type PaymentProviderAccountRepository struct {
	q *sqlc.Queries
}

func NewPaymentProviderAccountRepository(db *sql.DB) *PaymentProviderAccountRepository {
	return &PaymentProviderAccountRepository{q: sqlc.New(db)}
}

var _ payment.PaymentProviderAccountRepository = (*PaymentProviderAccountRepository)(nil)

func (r *PaymentProviderAccountRepository) Create(ctx context.Context, a *payment.PaymentProviderAccount) error {
	row, err := r.q.CreatePaymentProviderAccount(ctx, sqlc.CreatePaymentProviderAccountParams{
		ID:                 a.ID,
		PaymentProviderID:  a.PaymentProviderID,
		MemberID:           a.MemberID,
		ExternalAccountID:  a.ExternalAccountID,
	})
	if err != nil {
		return err
	}
	*a = toPaymentProviderAccountDomain(row)
	return nil
}

func (r *PaymentProviderAccountRepository) Get(ctx context.Context, id uuid.UUID) (*payment.PaymentProviderAccount, error) {
	row, err := r.q.GetPaymentProviderAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toPaymentProviderAccountDomain(row)
	return &a, nil
}

func (r *PaymentProviderAccountRepository) List(ctx context.Context) ([]*payment.PaymentProviderAccount, error) {
	rows, err := r.q.ListPaymentProviderAccounts(ctx)
	if err != nil {
		return nil, err
	}
	accounts := make([]*payment.PaymentProviderAccount, 0, len(rows))
	for _, row := range rows {
		a := toPaymentProviderAccountDomain(row)
		accounts = append(accounts, &a)
	}
	return accounts, nil
}

func (r *PaymentProviderAccountRepository) ListByPaymentProviderID(ctx context.Context, paymentProviderID uuid.UUID) ([]*payment.PaymentProviderAccount, error) {
	rows, err := r.q.ListPaymentProviderAccountsByPaymentProviderID(ctx, paymentProviderID)
	if err != nil {
		return nil, err
	}
	accounts := make([]*payment.PaymentProviderAccount, 0, len(rows))
	for _, row := range rows {
		a := toPaymentProviderAccountDomain(row)
		accounts = append(accounts, &a)
	}
	return accounts, nil
}

func (r *PaymentProviderAccountRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*payment.PaymentProviderAccount, error) {
	rows, err := r.q.ListPaymentProviderAccountsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	accounts := make([]*payment.PaymentProviderAccount, 0, len(rows))
	for _, row := range rows {
		a := toPaymentProviderAccountDomain(row)
		accounts = append(accounts, &a)
	}
	return accounts, nil
}

func (r *PaymentProviderAccountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePaymentProviderAccount(ctx, id)
}

func toPaymentProviderAccountDomain(row sqlc.PaymentProviderAccount) payment.PaymentProviderAccount {
	return payment.PaymentProviderAccount{
		ID:                 row.ID,
		PaymentProviderID:  row.PaymentProviderID,
		MemberID:           row.MemberID,
		ExternalAccountID:  row.ExternalAccountID,
	}
}
