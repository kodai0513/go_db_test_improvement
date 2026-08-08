package notification

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// EmailLog はメール送信ログを表す。
type EmailLog struct {
	ID       uuid.UUID
	MemberID uuid.UUID
	Subject  string
	SentAt   time.Time
}

// EmailLogRepository は email_logs テーブルへのアクセスを抽象化する。
type EmailLogRepository interface {
	Create(ctx context.Context, l *EmailLog) error
	Get(ctx context.Context, id uuid.UUID) (*EmailLog, error)
	List(ctx context.Context) ([]*EmailLog, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*EmailLog, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
