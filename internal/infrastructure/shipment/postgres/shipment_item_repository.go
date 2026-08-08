package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres/sqlc"
)

// ShipmentItemRepository は shipment.ShipmentItemRepository の実装。
type ShipmentItemRepository struct {
	q *sqlc.Queries
}

func NewShipmentItemRepository(db *sql.DB) *ShipmentItemRepository {
	return &ShipmentItemRepository{q: sqlc.New(db)}
}

var _ shipment.ShipmentItemRepository = (*ShipmentItemRepository)(nil)

func (r *ShipmentItemRepository) Create(ctx context.Context, si *shipment.ShipmentItem) error {
	row, err := r.q.CreateShipmentItem(ctx, sqlc.CreateShipmentItemParams{
		ID:          si.ID,
		ShipmentID:  si.ShipmentID,
		OrderItemID: si.OrderItemID,
		Quantity:    si.Quantity,
	})
	if err != nil {
		return err
	}
	*si = toShipmentItemDomain(row)
	return nil
}

func (r *ShipmentItemRepository) Get(ctx context.Context, id uuid.UUID) (*shipment.ShipmentItem, error) {
	row, err := r.q.GetShipmentItem(ctx, id)
	if err != nil {
		return nil, err
	}
	si := toShipmentItemDomain(row)
	return &si, nil
}

func (r *ShipmentItemRepository) List(ctx context.Context) ([]*shipment.ShipmentItem, error) {
	rows, err := r.q.ListShipmentItems(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*shipment.ShipmentItem, 0, len(rows))
	for _, row := range rows {
		si := toShipmentItemDomain(row)
		items = append(items, &si)
	}
	return items, nil
}

func (r *ShipmentItemRepository) ListByShipmentID(ctx context.Context, shipmentID uuid.UUID) ([]*shipment.ShipmentItem, error) {
	rows, err := r.q.ListShipmentItemsByShipmentID(ctx, shipmentID)
	if err != nil {
		return nil, err
	}
	items := make([]*shipment.ShipmentItem, 0, len(rows))
	for _, row := range rows {
		si := toShipmentItemDomain(row)
		items = append(items, &si)
	}
	return items, nil
}

func (r *ShipmentItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteShipmentItem(ctx, id)
}

func toShipmentItemDomain(row sqlc.ShipmentItem) shipment.ShipmentItem {
	return shipment.ShipmentItem{
		ID:          row.ID,
		ShipmentID:  row.ShipmentID,
		OrderItemID: row.OrderItemID,
		Quantity:    row.Quantity,
	}
}
