package review

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, r *Review) error
	Get(ctx context.Context, id uuid.UUID) (*Review, error)
	List(ctx context.Context) ([]*Review, error)
}
