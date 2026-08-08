package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres/sqlc"
)

// PaymentTransactionRepository は payment.PaymentTransactionRepository の実装。
type PaymentTransactionRepository struct {
	q *sqlc.Queries
}

func NewPaymentTransactionRepository(db *sql.DB) *PaymentTransactionRepository {
	return &PaymentTransactionRepository{q: sqlc.New(db)}
}

var _ payment.PaymentTransactionRepository = (*PaymentTransactionRepository)(nil)

func (r *PaymentTransactionRepository) Create(ctx context.Context, t *payment.PaymentTransaction) error {
	row, err := r.q.CreatePaymentTransaction(ctx, sqlc.CreatePaymentTransactionParams{
		ID:          t.ID,
		PaymentID:   t.PaymentID,
		Status:      t.Status,
		ProcessedAt: t.ProcessedAt,
	})
	if err != nil {
		return err
	}
	*t = toPaymentTransactionDomain(row)
	return nil
}

func (r *PaymentTransactionRepository) Get(ctx context.Context, id uuid.UUID) (*payment.PaymentTransaction, error) {
	row, err := r.q.GetPaymentTransaction(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toPaymentTransactionDomain(row)
	return &t, nil
}

func (r *PaymentTransactionRepository) List(ctx context.Context) ([]*payment.PaymentTransaction, error) {
	rows, err := r.q.ListPaymentTransactions(ctx)
	if err != nil {
		return nil, err
	}
	transactions := make([]*payment.PaymentTransaction, 0, len(rows))
	for _, row := range rows {
		t := toPaymentTransactionDomain(row)
		transactions = append(transactions, &t)
	}
	return transactions, nil
}

func (r *PaymentTransactionRepository) ListByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*payment.PaymentTransaction, error) {
	rows, err := r.q.ListPaymentTransactionsByPaymentID(ctx, paymentID)
	if err != nil {
		return nil, err
	}
	transactions := make([]*payment.PaymentTransaction, 0, len(rows))
	for _, row := range rows {
		t := toPaymentTransactionDomain(row)
		transactions = append(transactions, &t)
	}
	return transactions, nil
}

func (r *PaymentTransactionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePaymentTransaction(ctx, id)
}

func toPaymentTransactionDomain(row sqlc.PaymentTransaction) payment.PaymentTransaction {
	return payment.PaymentTransaction{
		ID:          row.ID,
		PaymentID:   row.PaymentID,
		Status:      row.Status,
		ProcessedAt: row.ProcessedAt,
	}
}
