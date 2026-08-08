package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres/sqlc"
)

// InstallmentPaymentRepository は payment.InstallmentPaymentRepository の実装。
type InstallmentPaymentRepository struct {
	q *sqlc.Queries
}

func NewInstallmentPaymentRepository(db *sql.DB) *InstallmentPaymentRepository {
	return &InstallmentPaymentRepository{q: sqlc.New(db)}
}

var _ payment.InstallmentPaymentRepository = (*InstallmentPaymentRepository)(nil)

func (r *InstallmentPaymentRepository) Create(ctx context.Context, p *payment.InstallmentPayment) error {
	row, err := r.q.CreateInstallmentPayment(ctx, sqlc.CreateInstallmentPaymentParams{
		ID:                p.ID,
		InstallmentPlanID: p.InstallmentPlanID,
		AmountYen:         p.AmountYen,
		DueAt:             p.DueAt,
		PaidAt:            p.PaidAt,
	})
	if err != nil {
		return err
	}
	*p = toInstallmentPaymentDomain(row)
	return nil
}

func (r *InstallmentPaymentRepository) Get(ctx context.Context, id uuid.UUID) (*payment.InstallmentPayment, error) {
	row, err := r.q.GetInstallmentPayment(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toInstallmentPaymentDomain(row)
	return &p, nil
}

func (r *InstallmentPaymentRepository) List(ctx context.Context) ([]*payment.InstallmentPayment, error) {
	rows, err := r.q.ListInstallmentPayments(ctx)
	if err != nil {
		return nil, err
	}
	payments := make([]*payment.InstallmentPayment, 0, len(rows))
	for _, row := range rows {
		p := toInstallmentPaymentDomain(row)
		payments = append(payments, &p)
	}
	return payments, nil
}

func (r *InstallmentPaymentRepository) ListByInstallmentPlanID(ctx context.Context, installmentPlanID uuid.UUID) ([]*payment.InstallmentPayment, error) {
	rows, err := r.q.ListInstallmentPaymentsByInstallmentPlanID(ctx, installmentPlanID)
	if err != nil {
		return nil, err
	}
	payments := make([]*payment.InstallmentPayment, 0, len(rows))
	for _, row := range rows {
		p := toInstallmentPaymentDomain(row)
		payments = append(payments, &p)
	}
	return payments, nil
}

func (r *InstallmentPaymentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteInstallmentPayment(ctx, id)
}

func toInstallmentPaymentDomain(row sqlc.InstallmentPayment) payment.InstallmentPayment {
	return payment.InstallmentPayment{
		ID:                row.ID,
		InstallmentPlanID: row.InstallmentPlanID,
		AmountYen:         row.AmountYen,
		DueAt:             row.DueAt,
		PaidAt:            row.PaidAt,
	}
}
