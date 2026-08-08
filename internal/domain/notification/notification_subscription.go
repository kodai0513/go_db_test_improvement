package notification

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// NotificationSubscription は会員のトピック購読を表す。
type NotificationSubscription struct {
	ID           uuid.UUID
	MemberID     uuid.UUID
	Topic        string
	SubscribedAt time.Time
}

// NotificationSubscriptionRepository は notification_subscriptions テーブルへのアクセスを抽象化する。
type NotificationSubscriptionRepository interface {
	Create(ctx context.Context, s *NotificationSubscription) error
	Get(ctx context.Context, id uuid.UUID) (*NotificationSubscription, error)
	List(ctx context.Context) ([]*NotificationSubscription, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*NotificationSubscription, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
