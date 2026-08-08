package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres/sqlc"
)

// NotificationLogRepository は notification.NotificationLogRepository の実装。
type NotificationLogRepository struct {
	q *sqlc.Queries
}

func NewNotificationLogRepository(db *sql.DB) *NotificationLogRepository {
	return &NotificationLogRepository{q: sqlc.New(db)}
}

var _ notification.NotificationLogRepository = (*NotificationLogRepository)(nil)

func (r *NotificationLogRepository) Create(ctx context.Context, l *notification.NotificationLog) error {
	row, err := r.q.CreateNotificationLog(ctx, sqlc.CreateNotificationLogParams{
		ID:             l.ID,
		NotificationID: l.NotificationID,
		Status:         l.Status,
		SentAt:         l.SentAt,
	})
	if err != nil {
		return err
	}
	*l = toNotificationLogDomain(row)
	return nil
}

func (r *NotificationLogRepository) Get(ctx context.Context, id uuid.UUID) (*notification.NotificationLog, error) {
	row, err := r.q.GetNotificationLog(ctx, id)
	if err != nil {
		return nil, err
	}
	l := toNotificationLogDomain(row)
	return &l, nil
}

func (r *NotificationLogRepository) List(ctx context.Context) ([]*notification.NotificationLog, error) {
	rows, err := r.q.ListNotificationLogs(ctx)
	if err != nil {
		return nil, err
	}
	logs := make([]*notification.NotificationLog, 0, len(rows))
	for _, row := range rows {
		l := toNotificationLogDomain(row)
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *NotificationLogRepository) ListByNotificationID(ctx context.Context, notificationID uuid.UUID) ([]*notification.NotificationLog, error) {
	rows, err := r.q.ListNotificationLogsByNotificationID(ctx, notificationID)
	if err != nil {
		return nil, err
	}
	logs := make([]*notification.NotificationLog, 0, len(rows))
	for _, row := range rows {
		l := toNotificationLogDomain(row)
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *NotificationLogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteNotificationLog(ctx, id)
}

func toNotificationLogDomain(row sqlc.NotificationLog) notification.NotificationLog {
	return notification.NotificationLog{
		ID:             row.ID,
		NotificationID: row.NotificationID,
		Status:         row.Status,
		SentAt:         row.SentAt,
	}
}
