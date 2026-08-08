package staff

import (
	"context"

	"github.com/google/uuid"
)

// Role はスタッフに割り当てられるロールを表す。
type Role struct {
	ID   uuid.UUID
	Name string
}

// RoleRepository は roles テーブルへのアクセスを抽象化する。
type RoleRepository interface {
	Create(ctx context.Context, r *Role) error
	Get(ctx context.Context, id uuid.UUID) (*Role, error)
	List(ctx context.Context) ([]*Role, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
