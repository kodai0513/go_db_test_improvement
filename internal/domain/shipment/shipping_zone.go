package shipment

import (
	"context"

	"github.com/google/uuid"
)

// ShippingZone は配送料金の地域区分を表す。
type ShippingZone struct {
	ID   uuid.UUID
	Name string
}

// ShippingZoneRepository は shipping_zones テーブルへのアクセスを抽象化する。
type ShippingZoneRepository interface {
	Create(ctx context.Context, sz *ShippingZone) error
	Get(ctx context.Context, id uuid.UUID) (*ShippingZone, error)
	List(ctx context.Context) ([]*ShippingZone, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
