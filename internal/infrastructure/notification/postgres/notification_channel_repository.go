package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres/sqlc"
)

// NotificationChannelRepository は notification.NotificationChannelRepository の実装。
type NotificationChannelRepository struct {
	q *sqlc.Queries
}

func NewNotificationChannelRepository(db *sql.DB) *NotificationChannelRepository {
	return &NotificationChannelRepository{q: sqlc.New(db)}
}

var _ notification.NotificationChannelRepository = (*NotificationChannelRepository)(nil)

func (r *NotificationChannelRepository) Create(ctx context.Context, c *notification.NotificationChannel) error {
	row, err := r.q.CreateNotificationChannel(ctx, sqlc.CreateNotificationChannelParams{
		ID:   c.ID,
		Name: c.Name,
	})
	if err != nil {
		return err
	}
	*c = toNotificationChannelDomain(row)
	return nil
}

func (r *NotificationChannelRepository) Get(ctx context.Context, id uuid.UUID) (*notification.NotificationChannel, error) {
	row, err := r.q.GetNotificationChannel(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toNotificationChannelDomain(row)
	return &c, nil
}

func (r *NotificationChannelRepository) List(ctx context.Context) ([]*notification.NotificationChannel, error) {
	rows, err := r.q.ListNotificationChannels(ctx)
	if err != nil {
		return nil, err
	}
	channels := make([]*notification.NotificationChannel, 0, len(rows))
	for _, row := range rows {
		c := toNotificationChannelDomain(row)
		channels = append(channels, &c)
	}
	return channels, nil
}

func (r *NotificationChannelRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteNotificationChannel(ctx, id)
}

func toNotificationChannelDomain(row sqlc.NotificationChannel) notification.NotificationChannel {
	return notification.NotificationChannel{
		ID:   row.ID,
		Name: row.Name,
	}
}
