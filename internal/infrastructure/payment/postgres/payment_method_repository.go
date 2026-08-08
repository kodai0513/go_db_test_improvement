package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres/sqlc"
)

// PaymentMethodRepository は payment.PaymentMethodRepository の実装。
type PaymentMethodRepository struct {
	q *sqlc.Queries
}

func NewPaymentMethodRepository(db *sql.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{q: sqlc.New(db)}
}

var _ payment.PaymentMethodRepository = (*PaymentMethodRepository)(nil)

func (r *PaymentMethodRepository) Create(ctx context.Context, p *payment.PaymentMethod) error {
	row, err := r.q.CreatePaymentMethod(ctx, sqlc.CreatePaymentMethodParams{
		ID:        p.ID,
		MemberID:  p.MemberID,
		Type:      p.Type,
		IsDefault: p.IsDefault,
		CreatedAt: p.CreatedAt,
	})
	if err != nil {
		return err
	}
	*p = toPaymentMethodDomain(row)
	return nil
}

func (r *PaymentMethodRepository) Get(ctx context.Context, id uuid.UUID) (*payment.PaymentMethod, error) {
	row, err := r.q.GetPaymentMethod(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPaymentMethodDomain(row)
	return &p, nil
}

func (r *PaymentMethodRepository) List(ctx context.Context) ([]*payment.PaymentMethod, error) {
	rows, err := r.q.ListPaymentMethods(ctx)
	if err != nil {
		return nil, err
	}
	methods := make([]*payment.PaymentMethod, 0, len(rows))
	for _, row := range rows {
		p := toPaymentMethodDomain(row)
		methods = append(methods, &p)
	}
	return methods, nil
}

func (r *PaymentMethodRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*payment.PaymentMethod, error) {
	rows, err := r.q.ListPaymentMethodsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	methods := make([]*payment.PaymentMethod, 0, len(rows))
	for _, row := range rows {
		p := toPaymentMethodDomain(row)
		methods = append(methods, &p)
	}
	return methods, nil
}

func (r *PaymentMethodRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePaymentMethod(ctx, id)
}

func toPaymentMethodDomain(row sqlc.PaymentMethod) payment.PaymentMethod {
	return payment.PaymentMethod{
		ID:        row.ID,
		MemberID:  row.MemberID,
		Type:      row.Type,
		IsDefault: row.IsDefault,
		CreatedAt: row.CreatedAt,
	}
}
