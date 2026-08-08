package payment

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, p *Payment) error
	Get(ctx context.Context, id uuid.UUID) (*Payment, error)
	List(ctx context.Context) ([]*Payment, error)
}
