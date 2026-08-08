package staff

import (
	"context"

	"github.com/google/uuid"
)

// Permission はロールに紐づく権限を表す。
type Permission struct {
	ID   uuid.UUID
	Name string
}

// PermissionRepository は permissions テーブルへのアクセスを抽象化する。
type PermissionRepository interface {
	Create(ctx context.Context, p *Permission) error
	Get(ctx context.Context, id uuid.UUID) (*Permission, error)
	List(ctx context.Context) ([]*Permission, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
