package staff

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// StaffSchedule はスタッフの勤務予定を表す。
type StaffSchedule struct {
	ID            uuid.UUID
	StaffMemberID uuid.UUID
	WorkDate      time.Time
	Shift         string
}

// StaffScheduleRepository は staff_schedules テーブルへのアクセスを抽象化する。
type StaffScheduleRepository interface {
	Create(ctx context.Context, s *StaffSchedule) error
	Get(ctx context.Context, id uuid.UUID) (*StaffSchedule, error)
	List(ctx context.Context) ([]*StaffSchedule, error)
	ListByStaffMemberID(ctx context.Context, staffMemberID uuid.UUID) ([]*StaffSchedule, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
