package review

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ReviewModerationLog はレビューに対するモデレーション操作の履歴を表す。
type ReviewModerationLog struct {
	ID        uuid.UUID
	ReviewID  uuid.UUID
	Action    string
	CreatedAt time.Time
}

// ReviewModerationLogRepository は review_moderation_logs テーブルへのアクセスを抽象化する。
type ReviewModerationLogRepository interface {
	Create(ctx context.Context, l *ReviewModerationLog) error
	Get(ctx context.Context, id uuid.UUID) (*ReviewModerationLog, error)
	List(ctx context.Context) ([]*ReviewModerationLog, error)
	ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*ReviewModerationLog, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
