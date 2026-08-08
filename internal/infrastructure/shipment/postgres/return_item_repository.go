package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres/sqlc"
)

// ReturnItemRepository は shipment.ReturnItemRepository の実装。
type ReturnItemRepository struct {
	q *sqlc.Queries
}

func NewReturnItemRepository(db *sql.DB) *ReturnItemRepository {
	return &ReturnItemRepository{q: sqlc.New(db)}
}

var _ shipment.ReturnItemRepository = (*ReturnItemRepository)(nil)

func (r *ReturnItemRepository) Create(ctx context.Context, ri *shipment.ReturnItem) error {
	row, err := r.q.CreateReturnItem(ctx, sqlc.CreateReturnItemParams{
		ID:              ri.ID,
		ReturnRequestID: ri.ReturnRequestID,
		OrderItemID:     ri.OrderItemID,
		Quantity:        ri.Quantity,
	})
	if err != nil {
		return err
	}
	*ri = toReturnItemDomain(row)
	return nil
}

func (r *ReturnItemRepository) Get(ctx context.Context, id uuid.UUID) (*shipment.ReturnItem, error) {
	row, err := r.q.GetReturnItem(ctx, id)
	if err != nil {
		return nil, err
	}
	ri := toReturnItemDomain(row)
	return &ri, nil
}

func (r *ReturnItemRepository) List(ctx context.Context) ([]*shipment.ReturnItem, error) {
	rows, err := r.q.ListReturnItems(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*shipment.ReturnItem, 0, len(rows))
	for _, row := range rows {
		ri := toReturnItemDomain(row)
		items = append(items, &ri)
	}
	return items, nil
}

func (r *ReturnItemRepository) ListByReturnRequestID(ctx context.Context, returnRequestID uuid.UUID) ([]*shipment.ReturnItem, error) {
	rows, err := r.q.ListReturnItemsByReturnRequestID(ctx, returnRequestID)
	if err != nil {
		return nil, err
	}
	items := make([]*shipment.ReturnItem, 0, len(rows))
	for _, row := range rows {
		ri := toReturnItemDomain(row)
		items = append(items, &ri)
	}
	return items, nil
}

func (r *ReturnItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReturnItem(ctx, id)
}

func toReturnItemDomain(row sqlc.ReturnItem) shipment.ReturnItem {
	return shipment.ReturnItem{
		ID:              row.ID,
		ReturnRequestID: row.ReturnRequestID,
		OrderItemID:     row.OrderItemID,
		Quantity:        row.Quantity,
	}
}
