package notification

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// PushDevice は会員に紐づくプッシュ通知デバイスを表す。
type PushDevice struct {
	ID           uuid.UUID
	MemberID     uuid.UUID
	DeviceToken  string
	RegisteredAt time.Time
}

// PushDeviceRepository は push_devices テーブルへのアクセスを抽象化する。
type PushDeviceRepository interface {
	Create(ctx context.Context, d *PushDevice) error
	Get(ctx context.Context, id uuid.UUID) (*PushDevice, error)
	List(ctx context.Context) ([]*PushDevice, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*PushDevice, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
