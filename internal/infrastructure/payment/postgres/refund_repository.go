package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres/sqlc"
)

// RefundRepository は payment.RefundRepository の実装。
type RefundRepository struct {
	q *sqlc.Queries
}

func NewRefundRepository(db *sql.DB) *RefundRepository {
	return &RefundRepository{q: sqlc.New(db)}
}

var _ payment.RefundRepository = (*RefundRepository)(nil)

func (r *RefundRepository) Create(ctx context.Context, ref *payment.Refund) error {
	row, err := r.q.CreateRefund(ctx, sqlc.CreateRefundParams{
		ID:         ref.ID,
		PaymentID:  ref.PaymentID,
		AmountYen:  ref.AmountYen,
		Reason:     ref.Reason,
		RefundedAt: ref.RefundedAt,
	})
	if err != nil {
		return err
	}
	*ref = toRefundDomain(row)
	return nil
}

func (r *RefundRepository) Get(ctx context.Context, id uuid.UUID) (*payment.Refund, error) {
	row, err := r.q.GetRefund(ctx, id)
	if err != nil {
		return nil, err
	}
	ref := toRefundDomain(row)
	return &ref, nil
}

func (r *RefundRepository) List(ctx context.Context) ([]*payment.Refund, error) {
	rows, err := r.q.ListRefunds(ctx)
	if err != nil {
		return nil, err
	}
	refunds := make([]*payment.Refund, 0, len(rows))
	for _, row := range rows {
		ref := toRefundDomain(row)
		refunds = append(refunds, &ref)
	}
	return refunds, nil
}

func (r *RefundRepository) ListByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*payment.Refund, error) {
	rows, err := r.q.ListRefundsByPaymentID(ctx, paymentID)
	if err != nil {
		return nil, err
	}
	refunds := make([]*payment.Refund, 0, len(rows))
	for _, row := range rows {
		ref := toRefundDomain(row)
		refunds = append(refunds, &ref)
	}
	return refunds, nil
}

func (r *RefundRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteRefund(ctx, id)
}

func toRefundDomain(row sqlc.Refund) payment.Refund {
	return payment.Refund{
		ID:         row.ID,
		PaymentID:  row.PaymentID,
		AmountYen:  row.AmountYen,
		Reason:     row.Reason,
		RefundedAt: row.RefundedAt,
	}
}
