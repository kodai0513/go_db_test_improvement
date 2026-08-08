package order

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, o *Order) error
	Get(ctx context.Context, id uuid.UUID) (*Order, error)
	List(ctx context.Context) ([]*Order, error)
}
