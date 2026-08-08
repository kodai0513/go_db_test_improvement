package shipment

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, s *Shipment) error
	Get(ctx context.Context, id uuid.UUID) (*Shipment, error)
	List(ctx context.Context) ([]*Shipment, error)
}
