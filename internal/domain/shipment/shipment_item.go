package shipment

import (
	"context"

	"github.com/google/uuid"
)

// ShipmentItem は出荷に含まれる商品明細を表す。
type ShipmentItem struct {
	ID          uuid.UUID
	ShipmentID  uuid.UUID
	OrderItemID uuid.UUID
	Quantity    int32
}

// ShipmentItemRepository は shipment_items テーブルへのアクセスを抽象化する。
type ShipmentItemRepository interface {
	Create(ctx context.Context, si *ShipmentItem) error
	Get(ctx context.Context, id uuid.UUID) (*ShipmentItem, error)
	List(ctx context.Context) ([]*ShipmentItem, error)
	ListByShipmentID(ctx context.Context, shipmentID uuid.UUID) ([]*ShipmentItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
