package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres/sqlc"
)

// ContentTranslationRepository は content.ContentTranslationRepository の実装。
type ContentTranslationRepository struct {
	q *sqlc.Queries
}

func NewContentTranslationRepository(db *sql.DB) *ContentTranslationRepository {
	return &ContentTranslationRepository{q: sqlc.New(db)}
}

var _ content.ContentTranslationRepository = (*ContentTranslationRepository)(nil)

func (r *ContentTranslationRepository) Create(ctx context.Context, t *content.ContentTranslation) error {
	row, err := r.q.CreateContentTranslation(ctx, sqlc.CreateContentTranslationParams{
		ID:             t.ID,
		PageID:         t.PageID,
		Locale:         t.Locale,
		TranslatedBody: t.TranslatedBody,
	})
	if err != nil {
		return err
	}
	*t = toContentTranslationDomain(row)
	return nil
}

func (r *ContentTranslationRepository) Get(ctx context.Context, id uuid.UUID) (*content.ContentTranslation, error) {
	row, err := r.q.GetContentTranslation(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toContentTranslationDomain(row)
	return &t, nil
}

func (r *ContentTranslationRepository) List(ctx context.Context) ([]*content.ContentTranslation, error) {
	rows, err := r.q.ListContentTranslations(ctx)
	if err != nil {
		return nil, err
	}
	translations := make([]*content.ContentTranslation, 0, len(rows))
	for _, row := range rows {
		t := toContentTranslationDomain(row)
		translations = append(translations, &t)
	}
	return translations, nil
}

func (r *ContentTranslationRepository) ListByPageID(ctx context.Context, pageID uuid.UUID) ([]*content.ContentTranslation, error) {
	rows, err := r.q.ListContentTranslationsByPageID(ctx, pageID)
	if err != nil {
		return nil, err
	}
	translations := make([]*content.ContentTranslation, 0, len(rows))
	for _, row := range rows {
		t := toContentTranslationDomain(row)
		translations = append(translations, &t)
	}
	return translations, nil
}

func (r *ContentTranslationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteContentTranslation(ctx, id)
}

func toContentTranslationDomain(row sqlc.ContentTranslation) content.ContentTranslation {
	return content.ContentTranslation{
		ID:             row.ID,
		PageID:         row.PageID,
		Locale:         row.Locale,
		TranslatedBody: row.TranslatedBody,
	}
}
