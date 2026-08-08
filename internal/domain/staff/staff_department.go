package staff

import (
	"context"

	"github.com/google/uuid"
)

// StaffDepartment はスタッフと部署の紐付けを表す。
type StaffDepartment struct {
	ID            uuid.UUID
	StaffMemberID uuid.UUID
	DepartmentID  uuid.UUID
}

// StaffDepartmentRepository は staff_departments テーブルへのアクセスを抽象化する。
type StaffDepartmentRepository interface {
	Create(ctx context.Context, sd *StaffDepartment) error
	Get(ctx context.Context, id uuid.UUID) (*StaffDepartment, error)
	List(ctx context.Context) ([]*StaffDepartment, error)
	ListByStaffMemberID(ctx context.Context, staffMemberID uuid.UUID) ([]*StaffDepartment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
