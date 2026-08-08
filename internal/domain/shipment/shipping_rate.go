package shipment

import (
	"context"

	"github.com/google/uuid"
)

// ShippingRate は配送業者ごとの配送料金を表す。
type ShippingRate struct {
	ID        uuid.UUID
	CarrierID uuid.UUID
	ZoneName  string
	PriceYen  int32
}

// ShippingRateRepository は shipping_rates テーブルへのアクセスを抽象化する。
type ShippingRateRepository interface {
	Create(ctx context.Context, sr *ShippingRate) error
	Get(ctx context.Context, id uuid.UUID) (*ShippingRate, error)
	List(ctx context.Context) ([]*ShippingRate, error)
	ListByCarrierID(ctx context.Context, carrierID uuid.UUID) ([]*ShippingRate, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
