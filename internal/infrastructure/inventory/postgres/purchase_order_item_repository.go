package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres/sqlc"
)

// PurchaseOrderItemRepository は inventory.PurchaseOrderItemRepository の実装。
type PurchaseOrderItemRepository struct {
	q *sqlc.Queries
}

func NewPurchaseOrderItemRepository(db *sql.DB) *PurchaseOrderItemRepository {
	return &PurchaseOrderItemRepository{q: sqlc.New(db)}
}

var _ inventory.PurchaseOrderItemRepository = (*PurchaseOrderItemRepository)(nil)

func (r *PurchaseOrderItemRepository) Create(ctx context.Context, i *inventory.PurchaseOrderItem) error {
	row, err := r.q.CreatePurchaseOrderItem(ctx, sqlc.CreatePurchaseOrderItemParams{
		ID:              i.ID,
		PurchaseOrderID: i.PurchaseOrderID,
		ProductID:       i.ProductID,
		Quantity:        i.Quantity,
		UnitPriceYen:    i.UnitPriceYen,
	})
	if err != nil {
		return err
	}
	*i = toPurchaseOrderItemDomain(row)
	return nil
}

func (r *PurchaseOrderItemRepository) Get(ctx context.Context, id uuid.UUID) (*inventory.PurchaseOrderItem, error) {
	row, err := r.q.GetPurchaseOrderItem(ctx, id)
	if err != nil {
		return nil, err
	}
	i := toPurchaseOrderItemDomain(row)
	return &i, nil
}

func (r *PurchaseOrderItemRepository) List(ctx context.Context) ([]*inventory.PurchaseOrderItem, error) {
	rows, err := r.q.ListPurchaseOrderItems(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*inventory.PurchaseOrderItem, 0, len(rows))
	for _, row := range rows {
		i := toPurchaseOrderItemDomain(row)
		items = append(items, &i)
	}
	return items, nil
}

func (r *PurchaseOrderItemRepository) ListByPurchaseOrderID(ctx context.Context, purchaseOrderID uuid.UUID) ([]*inventory.PurchaseOrderItem, error) {
	rows, err := r.q.ListPurchaseOrderItemsByPurchaseOrderID(ctx, purchaseOrderID)
	if err != nil {
		return nil, err
	}
	items := make([]*inventory.PurchaseOrderItem, 0, len(rows))
	for _, row := range rows {
		i := toPurchaseOrderItemDomain(row)
		items = append(items, &i)
	}
	return items, nil
}

func (r *PurchaseOrderItemRepository) ListByProductID(ctx context.Context, productID uuid.UUID) ([]*inventory.PurchaseOrderItem, error) {
	rows, err := r.q.ListPurchaseOrderItemsByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	items := make([]*inventory.PurchaseOrderItem, 0, len(rows))
	for _, row := range rows {
		i := toPurchaseOrderItemDomain(row)
		items = append(items, &i)
	}
	return items, nil
}

func (r *PurchaseOrderItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePurchaseOrderItem(ctx, id)
}

func toPurchaseOrderItemDomain(row sqlc.PurchaseOrderItem) inventory.PurchaseOrderItem {
	return inventory.PurchaseOrderItem{
		ID:              row.ID,
		PurchaseOrderID: row.PurchaseOrderID,
		ProductID:       row.ProductID,
		Quantity:        row.Quantity,
		UnitPriceYen:    row.UnitPriceYen,
	}
}
