package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres/sqlc"
)

// InstallmentPlanRepository は payment.InstallmentPlanRepository の実装。
type InstallmentPlanRepository struct {
	q *sqlc.Queries
}

func NewInstallmentPlanRepository(db *sql.DB) *InstallmentPlanRepository {
	return &InstallmentPlanRepository{q: sqlc.New(db)}
}

var _ payment.InstallmentPlanRepository = (*InstallmentPlanRepository)(nil)

func (r *InstallmentPlanRepository) Create(ctx context.Context, p *payment.InstallmentPlan) error {
	row, err := r.q.CreateInstallmentPlan(ctx, sqlc.CreateInstallmentPlanParams{
		ID:                   p.ID,
		PaymentID:            p.PaymentID,
		NumberOfInstallments: p.NumberOfInstallments,
	})
	if err != nil {
		return err
	}
	*p = toInstallmentPlanDomain(row)
	return nil
}

func (r *InstallmentPlanRepository) Get(ctx context.Context, id uuid.UUID) (*payment.InstallmentPlan, error) {
	row, err := r.q.GetInstallmentPlan(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toInstallmentPlanDomain(row)
	return &p, nil
}

func (r *InstallmentPlanRepository) List(ctx context.Context) ([]*payment.InstallmentPlan, error) {
	rows, err := r.q.ListInstallmentPlans(ctx)
	if err != nil {
		return nil, err
	}
	plans := make([]*payment.InstallmentPlan, 0, len(rows))
	for _, row := range rows {
		p := toInstallmentPlanDomain(row)
		plans = append(plans, &p)
	}
	return plans, nil
}

func (r *InstallmentPlanRepository) ListByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*payment.InstallmentPlan, error) {
	rows, err := r.q.ListInstallmentPlansByPaymentID(ctx, paymentID)
	if err != nil {
		return nil, err
	}
	plans := make([]*payment.InstallmentPlan, 0, len(rows))
	for _, row := range rows {
		p := toInstallmentPlanDomain(row)
		plans = append(plans, &p)
	}
	return plans, nil
}

func (r *InstallmentPlanRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteInstallmentPlan(ctx, id)
}

func toInstallmentPlanDomain(row sqlc.InstallmentPlan) payment.InstallmentPlan {
	return payment.InstallmentPlan{
		ID:                   row.ID,
		PaymentID:            row.PaymentID,
		NumberOfInstallments: row.NumberOfInstallments,
	}
}
