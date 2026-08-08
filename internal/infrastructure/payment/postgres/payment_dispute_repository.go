package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres/sqlc"
)

// PaymentDisputeRepository は payment.PaymentDisputeRepository の実装。
type PaymentDisputeRepository struct {
	q *sqlc.Queries
}

func NewPaymentDisputeRepository(db *sql.DB) *PaymentDisputeRepository {
	return &PaymentDisputeRepository{q: sqlc.New(db)}
}

var _ payment.PaymentDisputeRepository = (*PaymentDisputeRepository)(nil)

func (r *PaymentDisputeRepository) Create(ctx context.Context, d *payment.PaymentDispute) error {
	row, err := r.q.CreatePaymentDispute(ctx, sqlc.CreatePaymentDisputeParams{
		ID:        d.ID,
		PaymentID: d.PaymentID,
		Reason:    d.Reason,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
	})
	if err != nil {
		return err
	}
	*d = toPaymentDisputeDomain(row)
	return nil
}

func (r *PaymentDisputeRepository) Get(ctx context.Context, id uuid.UUID) (*payment.PaymentDispute, error) {
	row, err := r.q.GetPaymentDispute(ctx, id)
	if err != nil {
		return nil, err
	}
	d := toPaymentDisputeDomain(row)
	return &d, nil
}

func (r *PaymentDisputeRepository) List(ctx context.Context) ([]*payment.PaymentDispute, error) {
	rows, err := r.q.ListPaymentDisputes(ctx)
	if err != nil {
		return nil, err
	}
	disputes := make([]*payment.PaymentDispute, 0, len(rows))
	for _, row := range rows {
		d := toPaymentDisputeDomain(row)
		disputes = append(disputes, &d)
	}
	return disputes, nil
}

func (r *PaymentDisputeRepository) ListByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*payment.PaymentDispute, error) {
	rows, err := r.q.ListPaymentDisputesByPaymentID(ctx, paymentID)
	if err != nil {
		return nil, err
	}
	disputes := make([]*payment.PaymentDispute, 0, len(rows))
	for _, row := range rows {
		d := toPaymentDisputeDomain(row)
		disputes = append(disputes, &d)
	}
	return disputes, nil
}

func (r *PaymentDisputeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePaymentDispute(ctx, id)
}

func toPaymentDisputeDomain(row sqlc.PaymentDispute) payment.PaymentDispute {
	return payment.PaymentDispute{
		ID:        row.ID,
		PaymentID: row.PaymentID,
		Reason:    row.Reason,
		Status:    row.Status,
		CreatedAt: row.CreatedAt,
	}
}
