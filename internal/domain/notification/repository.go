package notification

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, n *Notification) error
	Get(ctx context.Context, id uuid.UUID) (*Notification, error)
	List(ctx context.Context) ([]*Notification, error)
}
