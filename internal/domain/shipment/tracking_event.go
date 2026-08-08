package shipment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// TrackingEvent は出荷の追跡イベントを表す。
type TrackingEvent struct {
	ID         uuid.UUID
	ShipmentID uuid.UUID
	Status     string
	OccurredAt time.Time
}

// TrackingEventRepository は tracking_events テーブルへのアクセスを抽象化する。
type TrackingEventRepository interface {
	Create(ctx context.Context, te *TrackingEvent) error
	Get(ctx context.Context, id uuid.UUID) (*TrackingEvent, error)
	List(ctx context.Context) ([]*TrackingEvent, error)
	ListByShipmentID(ctx context.Context, shipmentID uuid.UUID) ([]*TrackingEvent, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
