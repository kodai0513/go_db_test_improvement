package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres/sqlc"
)

// PaymentProviderRepository は payment.PaymentProviderRepository の実装。
type PaymentProviderRepository struct {
	q *sqlc.Queries
}

func NewPaymentProviderRepository(db *sql.DB) *PaymentProviderRepository {
	return &PaymentProviderRepository{q: sqlc.New(db)}
}

var _ payment.PaymentProviderRepository = (*PaymentProviderRepository)(nil)

func (r *PaymentProviderRepository) Create(ctx context.Context, p *payment.PaymentProvider) error {
	row, err := r.q.CreatePaymentProvider(ctx, sqlc.CreatePaymentProviderParams{
		ID:   p.ID,
		Name: p.Name,
	})
	if err != nil {
		return err
	}
	*p = toPaymentProviderDomain(row)
	return nil
}

func (r *PaymentProviderRepository) Get(ctx context.Context, id uuid.UUID) (*payment.PaymentProvider, error) {
	row, err := r.q.GetPaymentProvider(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPaymentProviderDomain(row)
	return &p, nil
}

func (r *PaymentProviderRepository) List(ctx context.Context) ([]*payment.PaymentProvider, error) {
	rows, err := r.q.ListPaymentProviders(ctx)
	if err != nil {
		return nil, err
	}
	providers := make([]*payment.PaymentProvider, 0, len(rows))
	for _, row := range rows {
		p := toPaymentProviderDomain(row)
		providers = append(providers, &p)
	}
	return providers, nil
}

func (r *PaymentProviderRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePaymentProvider(ctx, id)
}

func toPaymentProviderDomain(row sqlc.PaymentProvider) payment.PaymentProvider {
	return payment.PaymentProvider{
		ID:   row.ID,
		Name: row.Name,
	}
}
