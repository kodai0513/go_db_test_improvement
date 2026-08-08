package staff

import (
	"context"

	"github.com/google/uuid"
)

// StaffRole はスタッフとロールの紐付けを表す。
type StaffRole struct {
	ID            uuid.UUID
	StaffMemberID uuid.UUID
	RoleID        uuid.UUID
}

// StaffRoleRepository は staff_roles テーブルへのアクセスを抽象化する。
type StaffRoleRepository interface {
	Create(ctx context.Context, sr *StaffRole) error
	Get(ctx context.Context, id uuid.UUID) (*StaffRole, error)
	List(ctx context.Context) ([]*StaffRole, error)
	ListByStaffMemberID(ctx context.Context, staffMemberID uuid.UUID) ([]*StaffRole, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
