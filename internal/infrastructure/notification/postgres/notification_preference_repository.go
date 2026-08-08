package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres/sqlc"
)

// NotificationPreferenceRepository は notification.NotificationPreferenceRepository の実装。
type NotificationPreferenceRepository struct {
	q *sqlc.Queries
}

func NewNotificationPreferenceRepository(db *sql.DB) *NotificationPreferenceRepository {
	return &NotificationPreferenceRepository{q: sqlc.New(db)}
}

var _ notification.NotificationPreferenceRepository = (*NotificationPreferenceRepository)(nil)

func (r *NotificationPreferenceRepository) Create(ctx context.Context, p *notification.NotificationPreference) error {
	row, err := r.q.CreateNotificationPreference(ctx, sqlc.CreateNotificationPreferenceParams{
		ID:        p.ID,
		MemberID:  p.MemberID,
		ChannelID: p.ChannelID,
		Enabled:   p.Enabled,
	})
	if err != nil {
		return err
	}
	*p = toNotificationPreferenceDomain(row)
	return nil
}

func (r *NotificationPreferenceRepository) Get(ctx context.Context, id uuid.UUID) (*notification.NotificationPreference, error) {
	row, err := r.q.GetNotificationPreference(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toNotificationPreferenceDomain(row)
	return &p, nil
}

func (r *NotificationPreferenceRepository) List(ctx context.Context) ([]*notification.NotificationPreference, error) {
	rows, err := r.q.ListNotificationPreferences(ctx)
	if err != nil {
		return nil, err
	}
	prefs := make([]*notification.NotificationPreference, 0, len(rows))
	for _, row := range rows {
		p := toNotificationPreferenceDomain(row)
		prefs = append(prefs, &p)
	}
	return prefs, nil
}

func (r *NotificationPreferenceRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*notification.NotificationPreference, error) {
	rows, err := r.q.ListNotificationPreferencesByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	prefs := make([]*notification.NotificationPreference, 0, len(rows))
	for _, row := range rows {
		p := toNotificationPreferenceDomain(row)
		prefs = append(prefs, &p)
	}
	return prefs, nil
}

func (r *NotificationPreferenceRepository) ListByChannelID(ctx context.Context, channelID uuid.UUID) ([]*notification.NotificationPreference, error) {
	rows, err := r.q.ListNotificationPreferencesByChannelID(ctx, channelID)
	if err != nil {
		return nil, err
	}
	prefs := make([]*notification.NotificationPreference, 0, len(rows))
	for _, row := range rows {
		p := toNotificationPreferenceDomain(row)
		prefs = append(prefs, &p)
	}
	return prefs, nil
}

func (r *NotificationPreferenceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteNotificationPreference(ctx, id)
}

func toNotificationPreferenceDomain(row sqlc.NotificationPreference) notification.NotificationPreference {
	return notification.NotificationPreference{
		ID:        row.ID,
		MemberID:  row.MemberID,
		ChannelID: row.ChannelID,
		Enabled:   row.Enabled,
	}
}
