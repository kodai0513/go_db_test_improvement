package member

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, m *Member) error
	Get(ctx context.Context, id uuid.UUID) (*Member, error)
	List(ctx context.Context) ([]*Member, error)
}
