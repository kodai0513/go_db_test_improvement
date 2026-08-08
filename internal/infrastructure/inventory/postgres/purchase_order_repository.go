package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres/sqlc"
)

// PurchaseOrderRepository は inventory.PurchaseOrderRepository の実装。
type PurchaseOrderRepository struct {
	q *sqlc.Queries
}

func NewPurchaseOrderRepository(db *sql.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{q: sqlc.New(db)}
}

var _ inventory.PurchaseOrderRepository = (*PurchaseOrderRepository)(nil)

func (r *PurchaseOrderRepository) Create(ctx context.Context, p *inventory.PurchaseOrder) error {
	row, err := r.q.CreatePurchaseOrder(ctx, sqlc.CreatePurchaseOrderParams{
		ID:         p.ID,
		SupplierID: p.SupplierID,
		Status:     p.Status,
		CreatedAt:  p.CreatedAt,
	})
	if err != nil {
		return err
	}
	*p = toPurchaseOrderDomain(row)
	return nil
}

func (r *PurchaseOrderRepository) Get(ctx context.Context, id uuid.UUID) (*inventory.PurchaseOrder, error) {
	row, err := r.q.GetPurchaseOrder(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPurchaseOrderDomain(row)
	return &p, nil
}

func (r *PurchaseOrderRepository) List(ctx context.Context) ([]*inventory.PurchaseOrder, error) {
	rows, err := r.q.ListPurchaseOrders(ctx)
	if err != nil {
		return nil, err
	}
	orders := make([]*inventory.PurchaseOrder, 0, len(rows))
	for _, row := range rows {
		p := toPurchaseOrderDomain(row)
		orders = append(orders, &p)
	}
	return orders, nil
}

func (r *PurchaseOrderRepository) ListBySupplierID(ctx context.Context, supplierID uuid.UUID) ([]*inventory.PurchaseOrder, error) {
	rows, err := r.q.ListPurchaseOrdersBySupplierID(ctx, supplierID)
	if err != nil {
		return nil, err
	}
	orders := make([]*inventory.PurchaseOrder, 0, len(rows))
	for _, row := range rows {
		p := toPurchaseOrderDomain(row)
		orders = append(orders, &p)
	}
	return orders, nil
}

func (r *PurchaseOrderRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePurchaseOrder(ctx, id)
}

func toPurchaseOrderDomain(row sqlc.PurchaseOrder) inventory.PurchaseOrder {
	return inventory.PurchaseOrder{
		ID:         row.ID,
		SupplierID: row.SupplierID,
		Status:     row.Status,
		CreatedAt:  row.CreatedAt,
	}
}
