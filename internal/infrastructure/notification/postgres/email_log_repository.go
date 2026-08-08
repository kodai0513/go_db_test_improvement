package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres/sqlc"
)

// EmailLogRepository は notification.EmailLogRepository の実装。
type EmailLogRepository struct {
	q *sqlc.Queries
}

func NewEmailLogRepository(db *sql.DB) *EmailLogRepository {
	return &EmailLogRepository{q: sqlc.New(db)}
}

var _ notification.EmailLogRepository = (*EmailLogRepository)(nil)

func (r *EmailLogRepository) Create(ctx context.Context, l *notification.EmailLog) error {
	row, err := r.q.CreateEmailLog(ctx, sqlc.CreateEmailLogParams{
		ID:       l.ID,
		MemberID: l.MemberID,
		Subject:  l.Subject,
		SentAt:   l.SentAt,
	})
	if err != nil {
		return err
	}
	*l = toEmailLogDomain(row)
	return nil
}

func (r *EmailLogRepository) Get(ctx context.Context, id uuid.UUID) (*notification.EmailLog, error) {
	row, err := r.q.GetEmailLog(ctx, id)
	if err != nil {
		return nil, err
	}
	l := toEmailLogDomain(row)
	return &l, nil
}

func (r *EmailLogRepository) List(ctx context.Context) ([]*notification.EmailLog, error) {
	rows, err := r.q.ListEmailLogs(ctx)
	if err != nil {
		return nil, err
	}
	logs := make([]*notification.EmailLog, 0, len(rows))
	for _, row := range rows {
		l := toEmailLogDomain(row)
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *EmailLogRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*notification.EmailLog, error) {
	rows, err := r.q.ListEmailLogsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	logs := make([]*notification.EmailLog, 0, len(rows))
	for _, row := range rows {
		l := toEmailLogDomain(row)
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *EmailLogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteEmailLog(ctx, id)
}

func toEmailLogDomain(row sqlc.EmailLog) notification.EmailLog {
	return notification.EmailLog{
		ID:       row.ID,
		MemberID: row.MemberID,
		Subject:  row.Subject,
		SentAt:   row.SentAt,
	}
}
