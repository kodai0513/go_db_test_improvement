package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// InstallmentPayment は分割払いプランに紐づく個々の支払回を表す。
// paid_at はDB上はNULL許容だが、Go側は未払いの場合ゼロ値のtime.Timeを書き込む運用とする。
type InstallmentPayment struct {
	ID                uuid.UUID
	InstallmentPlanID uuid.UUID
	AmountYen         int32
	DueAt             time.Time
	PaidAt            time.Time
}

// InstallmentPaymentRepository は installment_payments テーブルへのアクセスを抽象化する。
type InstallmentPaymentRepository interface {
	Create(ctx context.Context, p *InstallmentPayment) error
	Get(ctx context.Context, id uuid.UUID) (*InstallmentPayment, error)
	List(ctx context.Context) ([]*InstallmentPayment, error)
	ListByInstallmentPlanID(ctx context.Context, installmentPlanID uuid.UUID) ([]*InstallmentPayment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
