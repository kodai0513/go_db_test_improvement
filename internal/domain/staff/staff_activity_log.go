package staff

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// StaffActivityLog はスタッフの操作履歴を表す。
type StaffActivityLog struct {
	ID            uuid.UUID
	StaffMemberID uuid.UUID
	Action        string
	OccurredAt    time.Time
}

// StaffActivityLogRepository は staff_activity_logs テーブルへのアクセスを抽象化する。
type StaffActivityLogRepository interface {
	Create(ctx context.Context, l *StaffActivityLog) error
	Get(ctx context.Context, id uuid.UUID) (*StaffActivityLog, error)
	List(ctx context.Context) ([]*StaffActivityLog, error)
	ListByStaffMemberID(ctx context.Context, staffMemberID uuid.UUID) ([]*StaffActivityLog, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
