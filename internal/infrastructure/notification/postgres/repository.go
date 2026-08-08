package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres/sqlc"
)

// Repository implements notification.Repository backed by the
// sqlc-generated query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ notification.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, n *notification.Notification) error {
	row, err := r.q.CreateNotification(ctx, sqlc.CreateNotificationParams{
		ID:        n.ID,
		MemberID:  n.MemberID,
		Message:   n.Message,
		IsRead:    n.IsRead,
		CreatedAt: n.CreatedAt,
	})
	if err != nil {
		return err
	}
	*n = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*notification.Notification, error) {
	row, err := r.q.GetNotification(ctx, id)
	if err != nil {
		return nil, err
	}
	n := toDomain(row)
	return &n, nil
}

func (r *Repository) List(ctx context.Context) ([]*notification.Notification, error) {
	rows, err := r.q.ListNotifications(ctx)
	if err != nil {
		return nil, err
	}
	notifications := make([]*notification.Notification, 0, len(rows))
	for _, row := range rows {
		n := toDomain(row)
		notifications = append(notifications, &n)
	}
	return notifications, nil
}

func toDomain(row sqlc.Notification) notification.Notification {
	return notification.Notification{
		ID:        row.ID,
		MemberID:  row.MemberID,
		Message:   row.Message,
		IsRead:    row.IsRead,
		CreatedAt: row.CreatedAt,
	}
}
