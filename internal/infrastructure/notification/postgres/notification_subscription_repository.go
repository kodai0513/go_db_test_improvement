package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres/sqlc"
)

// NotificationSubscriptionRepository は notification.NotificationSubscriptionRepository の実装。
type NotificationSubscriptionRepository struct {
	q *sqlc.Queries
}

func NewNotificationSubscriptionRepository(db *sql.DB) *NotificationSubscriptionRepository {
	return &NotificationSubscriptionRepository{q: sqlc.New(db)}
}

var _ notification.NotificationSubscriptionRepository = (*NotificationSubscriptionRepository)(nil)

func (r *NotificationSubscriptionRepository) Create(ctx context.Context, s *notification.NotificationSubscription) error {
	row, err := r.q.CreateNotificationSubscription(ctx, sqlc.CreateNotificationSubscriptionParams{
		ID:           s.ID,
		MemberID:     s.MemberID,
		Topic:        s.Topic,
		SubscribedAt: s.SubscribedAt,
	})
	if err != nil {
		return err
	}
	*s = toNotificationSubscriptionDomain(row)
	return nil
}

func (r *NotificationSubscriptionRepository) Get(ctx context.Context, id uuid.UUID) (*notification.NotificationSubscription, error) {
	row, err := r.q.GetNotificationSubscription(ctx, id)
	if err != nil {
		return nil, err
	}
	s := toNotificationSubscriptionDomain(row)
	return &s, nil
}

func (r *NotificationSubscriptionRepository) List(ctx context.Context) ([]*notification.NotificationSubscription, error) {
	rows, err := r.q.ListNotificationSubscriptions(ctx)
	if err != nil {
		return nil, err
	}
	subs := make([]*notification.NotificationSubscription, 0, len(rows))
	for _, row := range rows {
		s := toNotificationSubscriptionDomain(row)
		subs = append(subs, &s)
	}
	return subs, nil
}

func (r *NotificationSubscriptionRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*notification.NotificationSubscription, error) {
	rows, err := r.q.ListNotificationSubscriptionsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	subs := make([]*notification.NotificationSubscription, 0, len(rows))
	for _, row := range rows {
		s := toNotificationSubscriptionDomain(row)
		subs = append(subs, &s)
	}
	return subs, nil
}

func (r *NotificationSubscriptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteNotificationSubscription(ctx, id)
}

func toNotificationSubscriptionDomain(row sqlc.NotificationSubscription) notification.NotificationSubscription {
	return notification.NotificationSubscription{
		ID:           row.ID,
		MemberID:     row.MemberID,
		Topic:        row.Topic,
		SubscribedAt: row.SubscribedAt,
	}
}
