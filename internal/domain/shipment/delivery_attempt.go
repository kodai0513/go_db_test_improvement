package shipment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// DeliveryAttempt は配送の試行結果を表す。
type DeliveryAttempt struct {
	ID          uuid.UUID
	ShipmentID  uuid.UUID
	AttemptedAt time.Time
	Succeeded   bool
}

// DeliveryAttemptRepository は delivery_attempts テーブルへのアクセスを抽象化する。
type DeliveryAttemptRepository interface {
	Create(ctx context.Context, da *DeliveryAttempt) error
	Get(ctx context.Context, id uuid.UUID) (*DeliveryAttempt, error)
	List(ctx context.Context) ([]*DeliveryAttempt, error)
	ListByShipmentID(ctx context.Context, shipmentID uuid.UUID) ([]*DeliveryAttempt, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
