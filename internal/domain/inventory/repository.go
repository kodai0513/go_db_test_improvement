package inventory

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, i *Inventory) error
	Get(ctx context.Context, id uuid.UUID) (*Inventory, error)
	List(ctx context.Context) ([]*Inventory, error)
}
