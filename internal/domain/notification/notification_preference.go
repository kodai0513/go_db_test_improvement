package notification

import (
	"context"

	"github.com/google/uuid"
)

// NotificationPreference は会員ごとの通知チャネル設定を表す。
type NotificationPreference struct {
	ID        uuid.UUID
	MemberID  uuid.UUID
	ChannelID uuid.UUID
	Enabled   bool
}

// NotificationPreferenceRepository は notification_preferences テーブルへのアクセスを抽象化する。
type NotificationPreferenceRepository interface {
	Create(ctx context.Context, p *NotificationPreference) error
	Get(ctx context.Context, id uuid.UUID) (*NotificationPreference, error)
	List(ctx context.Context) ([]*NotificationPreference, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*NotificationPreference, error)
	ListByChannelID(ctx context.Context, channelID uuid.UUID) ([]*NotificationPreference, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
