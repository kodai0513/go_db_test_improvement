package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres/sqlc"
)

// ReviewModerationLogRepository は review.ReviewModerationLogRepository の実装。
type ReviewModerationLogRepository struct {
	q *sqlc.Queries
}

func NewReviewModerationLogRepository(db *sql.DB) *ReviewModerationLogRepository {
	return &ReviewModerationLogRepository{q: sqlc.New(db)}
}

var _ review.ReviewModerationLogRepository = (*ReviewModerationLogRepository)(nil)

func (r *ReviewModerationLogRepository) Create(ctx context.Context, l *review.ReviewModerationLog) error {
	row, err := r.q.CreateReviewModerationLog(ctx, sqlc.CreateReviewModerationLogParams{
		ID:        l.ID,
		ReviewID:  l.ReviewID,
		Action:    l.Action,
		CreatedAt: l.CreatedAt,
	})
	if err != nil {
		return err
	}
	*l = toReviewModerationLogDomain(row)
	return nil
}

func (r *ReviewModerationLogRepository) Get(ctx context.Context, id uuid.UUID) (*review.ReviewModerationLog, error) {
	row, err := r.q.GetReviewModerationLog(ctx, id)
	if err != nil {
		return nil, err
	}
	l := toReviewModerationLogDomain(row)
	return &l, nil
}

func (r *ReviewModerationLogRepository) List(ctx context.Context) ([]*review.ReviewModerationLog, error) {
	rows, err := r.q.ListReviewModerationLogs(ctx)
	if err != nil {
		return nil, err
	}
	logs := make([]*review.ReviewModerationLog, 0, len(rows))
	for _, row := range rows {
		l := toReviewModerationLogDomain(row)
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *ReviewModerationLogRepository) ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*review.ReviewModerationLog, error) {
	rows, err := r.q.ListReviewModerationLogsByReviewID(ctx, reviewID)
	if err != nil {
		return nil, err
	}
	logs := make([]*review.ReviewModerationLog, 0, len(rows))
	for _, row := range rows {
		l := toReviewModerationLogDomain(row)
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *ReviewModerationLogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReviewModerationLog(ctx, id)
}

func toReviewModerationLogDomain(row sqlc.ReviewModerationLog) review.ReviewModerationLog {
	return review.ReviewModerationLog{
		ID:        row.ID,
		ReviewID:  row.ReviewID,
		Action:    row.Action,
		CreatedAt: row.CreatedAt,
	}
}
