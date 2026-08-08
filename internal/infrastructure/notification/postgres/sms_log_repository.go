package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres/sqlc"
)

// SMSLogRepository は notification.SMSLogRepository の実装。
type SMSLogRepository struct {
	q *sqlc.Queries
}

func NewSMSLogRepository(db *sql.DB) *SMSLogRepository {
	return &SMSLogRepository{q: sqlc.New(db)}
}

var _ notification.SMSLogRepository = (*SMSLogRepository)(nil)

func (r *SMSLogRepository) Create(ctx context.Context, l *notification.SMSLog) error {
	row, err := r.q.CreateSMSLog(ctx, sqlc.CreateSMSLogParams{
		ID:       l.ID,
		MemberID: l.MemberID,
		Body:     l.Body,
		SentAt:   l.SentAt,
	})
	if err != nil {
		return err
	}
	*l = toSMSLogDomain(row)
	return nil
}

func (r *SMSLogRepository) Get(ctx context.Context, id uuid.UUID) (*notification.SMSLog, error) {
	row, err := r.q.GetSMSLog(ctx, id)
	if err != nil {
		return nil, err
	}
	l := toSMSLogDomain(row)
	return &l, nil
}

func (r *SMSLogRepository) List(ctx context.Context) ([]*notification.SMSLog, error) {
	rows, err := r.q.ListSMSLogs(ctx)
	if err != nil {
		return nil, err
	}
	logs := make([]*notification.SMSLog, 0, len(rows))
	for _, row := range rows {
		l := toSMSLogDomain(row)
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *SMSLogRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*notification.SMSLog, error) {
	rows, err := r.q.ListSMSLogsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	logs := make([]*notification.SMSLog, 0, len(rows))
	for _, row := range rows {
		l := toSMSLogDomain(row)
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *SMSLogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteSMSLog(ctx, id)
}

func toSMSLogDomain(row sqlc.SMSLog) notification.SMSLog {
	return notification.SMSLog{
		ID:       row.ID,
		MemberID: row.MemberID,
		Body:     row.Body,
		SentAt:   row.SentAt,
	}
}
