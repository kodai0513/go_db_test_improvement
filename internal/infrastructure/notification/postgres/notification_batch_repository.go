package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres/sqlc"
)

// NotificationBatchRepository は notification.NotificationBatchRepository の実装。
type NotificationBatchRepository struct {
	q *sqlc.Queries
}

func NewNotificationBatchRepository(db *sql.DB) *NotificationBatchRepository {
	return &NotificationBatchRepository{q: sqlc.New(db)}
}

var _ notification.NotificationBatchRepository = (*NotificationBatchRepository)(nil)

func (r *NotificationBatchRepository) Create(ctx context.Context, b *notification.NotificationBatch) error {
	row, err := r.q.CreateNotificationBatch(ctx, sqlc.CreateNotificationBatchParams{
		ID:         b.ID,
		TemplateID: b.TemplateID,
		Status:     b.Status,
		CreatedAt:  b.CreatedAt,
	})
	if err != nil {
		return err
	}
	*b = toNotificationBatchDomain(row)
	return nil
}

func (r *NotificationBatchRepository) Get(ctx context.Context, id uuid.UUID) (*notification.NotificationBatch, error) {
	row, err := r.q.GetNotificationBatch(ctx, id)
	if err != nil {
		return nil, err
	}
	b := toNotificationBatchDomain(row)
	return &b, nil
}

func (r *NotificationBatchRepository) List(ctx context.Context) ([]*notification.NotificationBatch, error) {
	rows, err := r.q.ListNotificationBatches(ctx)
	if err != nil {
		return nil, err
	}
	batches := make([]*notification.NotificationBatch, 0, len(rows))
	for _, row := range rows {
		b := toNotificationBatchDomain(row)
		batches = append(batches, &b)
	}
	return batches, nil
}

func (r *NotificationBatchRepository) ListByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*notification.NotificationBatch, error) {
	rows, err := r.q.ListNotificationBatchesByTemplateID(ctx, templateID)
	if err != nil {
		return nil, err
	}
	batches := make([]*notification.NotificationBatch, 0, len(rows))
	for _, row := range rows {
		b := toNotificationBatchDomain(row)
		batches = append(batches, &b)
	}
	return batches, nil
}

func (r *NotificationBatchRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteNotificationBatch(ctx, id)
}

func toNotificationBatchDomain(row sqlc.NotificationBatch) notification.NotificationBatch {
	return notification.NotificationBatch{
		ID:         row.ID,
		TemplateID: row.TemplateID,
		Status:     row.Status,
		CreatedAt:  row.CreatedAt,
	}
}
