package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres/sqlc"
)

// InvoiceRepository は payment.InvoiceRepository の実装。
type InvoiceRepository struct {
	q *sqlc.Queries
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{q: sqlc.New(db)}
}

var _ payment.InvoiceRepository = (*InvoiceRepository)(nil)

func (r *InvoiceRepository) Create(ctx context.Context, i *payment.Invoice) error {
	row, err := r.q.CreateInvoice(ctx, sqlc.CreateInvoiceParams{
		ID:        i.ID,
		OrderID:   i.OrderID,
		AmountYen: i.AmountYen,
		IssuedAt:  i.IssuedAt,
	})
	if err != nil {
		return err
	}
	*i = toInvoiceDomain(row)
	return nil
}

func (r *InvoiceRepository) Get(ctx context.Context, id uuid.UUID) (*payment.Invoice, error) {
	row, err := r.q.GetInvoice(ctx, id)
	if err != nil {
		return nil, err
	}
	i := toInvoiceDomain(row)
	return &i, nil
}

func (r *InvoiceRepository) List(ctx context.Context) ([]*payment.Invoice, error) {
	rows, err := r.q.ListInvoices(ctx)
	if err != nil {
		return nil, err
	}
	invoices := make([]*payment.Invoice, 0, len(rows))
	for _, row := range rows {
		i := toInvoiceDomain(row)
		invoices = append(invoices, &i)
	}
	return invoices, nil
}

func (r *InvoiceRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*payment.Invoice, error) {
	rows, err := r.q.ListInvoicesByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	invoices := make([]*payment.Invoice, 0, len(rows))
	for _, row := range rows {
		i := toInvoiceDomain(row)
		invoices = append(invoices, &i)
	}
	return invoices, nil
}

func (r *InvoiceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteInvoice(ctx, id)
}

func toInvoiceDomain(row sqlc.Invoice) payment.Invoice {
	return payment.Invoice{
		ID:        row.ID,
		OrderID:   row.OrderID,
		AmountYen: row.AmountYen,
		IssuedAt:  row.IssuedAt,
	}
}
