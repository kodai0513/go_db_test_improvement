package staff

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// StaffMember はスタッフ(従業員)を表す。
type StaffMember struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
}

// StaffMemberRepository は staff_members テーブルへのアクセスを抽象化する。
type StaffMemberRepository interface {
	Create(ctx context.Context, m *StaffMember) error
	Get(ctx context.Context, id uuid.UUID) (*StaffMember, error)
	List(ctx context.Context) ([]*StaffMember, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
