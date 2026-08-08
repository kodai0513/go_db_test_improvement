package product

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, p *Product) error
	Get(ctx context.Context, id uuid.UUID) (*Product, error)
	List(ctx context.Context) ([]*Product, error)
}
