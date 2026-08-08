package support

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// SatisfactionSurvey はチケットに対する満足度アンケートを表す。
type SatisfactionSurvey struct {
	ID          uuid.UUID
	TicketID    uuid.UUID
	Score       int32
	Comment     string
	SubmittedAt time.Time
}

// SatisfactionSurveyRepository は satisfaction_surveys テーブルへのアクセスを抽象化する。
type SatisfactionSurveyRepository interface {
	Create(ctx context.Context, s *SatisfactionSurvey) error
	Get(ctx context.Context, id uuid.UUID) (*SatisfactionSurvey, error)
	List(ctx context.Context) ([]*SatisfactionSurvey, error)
	ListByTicketID(ctx context.Context, ticketID uuid.UUID) ([]*SatisfactionSurvey, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
