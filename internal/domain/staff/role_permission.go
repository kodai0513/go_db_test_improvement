package staff

import (
	"context"

	"github.com/google/uuid"
)

// RolePermission はロールと権限の紐付けを表す。
type RolePermission struct {
	ID           uuid.UUID
	RoleID       uuid.UUID
	PermissionID uuid.UUID
}

// RolePermissionRepository は role_permissions テーブルへのアクセスを抽象化する。
type RolePermissionRepository interface {
	Create(ctx context.Context, rp *RolePermission) error
	Get(ctx context.Context, id uuid.UUID) (*RolePermission, error)
	List(ctx context.Context) ([]*RolePermission, error)
	ListByRoleID(ctx context.Context, roleID uuid.UUID) ([]*RolePermission, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
