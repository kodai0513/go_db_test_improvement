package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres/sqlc"
)

// NotificationTemplateRepository は notification.NotificationTemplateRepository の実装。
type NotificationTemplateRepository struct {
	q *sqlc.Queries
}

func NewNotificationTemplateRepository(db *sql.DB) *NotificationTemplateRepository {
	return &NotificationTemplateRepository{q: sqlc.New(db)}
}

var _ notification.NotificationTemplateRepository = (*NotificationTemplateRepository)(nil)

func (r *NotificationTemplateRepository) Create(ctx context.Context, t *notification.NotificationTemplate) error {
	row, err := r.q.CreateNotificationTemplate(ctx, sqlc.CreateNotificationTemplateParams{
		ID:   t.ID,
		Name: t.Name,
		Body: t.Body,
	})
	if err != nil {
		return err
	}
	*t = toNotificationTemplateDomain(row)
	return nil
}

func (r *NotificationTemplateRepository) Get(ctx context.Context, id uuid.UUID) (*notification.NotificationTemplate, error) {
	row, err := r.q.GetNotificationTemplate(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toNotificationTemplateDomain(row)
	return &t, nil
}

func (r *NotificationTemplateRepository) List(ctx context.Context) ([]*notification.NotificationTemplate, error) {
	rows, err := r.q.ListNotificationTemplates(ctx)
	if err != nil {
		return nil, err
	}
	templates := make([]*notification.NotificationTemplate, 0, len(rows))
	for _, row := range rows {
		t := toNotificationTemplateDomain(row)
		templates = append(templates, &t)
	}
	return templates, nil
}

func (r *NotificationTemplateRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteNotificationTemplate(ctx, id)
}

func toNotificationTemplateDomain(row sqlc.NotificationTemplate) notification.NotificationTemplate {
	return notification.NotificationTemplate{
		ID:   row.ID,
		Name: row.Name,
		Body: row.Body,
	}
}
