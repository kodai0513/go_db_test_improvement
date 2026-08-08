package notification

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// SMSLog はSMS送信ログを表す。
type SMSLog struct {
	ID       uuid.UUID
	MemberID uuid.UUID
	Body     string
	SentAt   time.Time
}

// SMSLogRepository は sms_logs テーブルへのアクセスを抽象化する。
type SMSLogRepository interface {
	Create(ctx context.Context, l *SMSLog) error
	Get(ctx context.Context, id uuid.UUID) (*SMSLog, error)
	List(ctx context.Context) ([]*SMSLog, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*SMSLog, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
