package shipment

import (
	"context"

	"github.com/google/uuid"
)

// ShippingAddress は出荷先住所を表す。
type ShippingAddress struct {
	ID         uuid.UUID
	ShipmentID uuid.UUID
	PostalCode string
	Address    string
}

// ShippingAddressRepository は shipping_addresses テーブルへのアクセスを抽象化する。
type ShippingAddressRepository interface {
	Create(ctx context.Context, sa *ShippingAddress) error
	Get(ctx context.Context, id uuid.UUID) (*ShippingAddress, error)
	List(ctx context.Context) ([]*ShippingAddress, error)
	ListByShipmentID(ctx context.Context, shipmentID uuid.UUID) ([]*ShippingAddress, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
