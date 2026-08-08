package coupon

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, c *Coupon) error
	Get(ctx context.Context, id uuid.UUID) (*Coupon, error)
	List(ctx context.Context) ([]*Coupon, error)
}
