package shipment

import (
	"context"

	"github.com/google/uuid"
)

// Carrier は配送業者を表す。
type Carrier struct {
	ID   uuid.UUID
	Name string
}

// CarrierRepository は carriers テーブルへのアクセスを抽象化する。
type CarrierRepository interface {
	Create(ctx context.Context, c *Carrier) error
	Get(ctx context.Context, id uuid.UUID) (*Carrier, error)
	List(ctx context.Context) ([]*Carrier, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
