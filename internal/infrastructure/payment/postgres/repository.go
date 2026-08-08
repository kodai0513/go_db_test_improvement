package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres/sqlc"
)

// Repository implements payment.Repository backed by the sqlc-generated
// query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ payment.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, p *payment.Payment) error {
	row, err := r.q.CreatePayment(ctx, sqlc.CreatePaymentParams{
		ID:        p.ID,
		OrderID:   p.OrderID,
		AmountYen: p.AmountYen,
		Method:    p.Method,
		PaidAt:    p.PaidAt,
	})
	if err != nil {
		return err
	}
	*p = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*payment.Payment, error) {
	row, err := r.q.GetPayment(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toDomain(row)
	return &p, nil
}

func (r *Repository) List(ctx context.Context) ([]*payment.Payment, error) {
	rows, err := r.q.ListPayments(ctx)
	if err != nil {
		return nil, err
	}
	payments := make([]*payment.Payment, 0, len(rows))
	for _, row := range rows {
		p := toDomain(row)
		payments = append(payments, &p)
	}
	return payments, nil
}

func toDomain(row sqlc.Payment) payment.Payment {
	return payment.Payment{
		ID:        row.ID,
		OrderID:   row.OrderID,
		AmountYen: row.AmountYen,
		Method:    row.Method,
		PaidAt:    row.PaidAt,
	}
}
