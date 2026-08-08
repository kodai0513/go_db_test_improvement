package payment

import (
	"context"

	"github.com/google/uuid"
)

// InstallmentPlan は決済に対する分割払いプランを表す。
type InstallmentPlan struct {
	ID                   uuid.UUID
	PaymentID            uuid.UUID
	NumberOfInstallments int32
}

// InstallmentPlanRepository は installment_plans テーブルへのアクセスを抽象化する。
type InstallmentPlanRepository interface {
	Create(ctx context.Context, p *InstallmentPlan) error
	Get(ctx context.Context, id uuid.UUID) (*InstallmentPlan, error)
	List(ctx context.Context) ([]*InstallmentPlan, error)
	ListByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*InstallmentPlan, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
