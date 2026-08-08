package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres/sqlc"
)

// FaqRepository は support.FaqRepository の実装。
type FaqRepository struct {
	q *sqlc.Queries
}

func NewFaqRepository(db *sql.DB) *FaqRepository {
	return &FaqRepository{q: sqlc.New(db)}
}

var _ support.FaqRepository = (*FaqRepository)(nil)

func (r *FaqRepository) Create(ctx context.Context, f *support.Faq) error {
	row, err := r.q.CreateFaq(ctx, sqlc.CreateFaqParams{
		ID:        f.ID,
		Question:  f.Question,
		Answer:    f.Answer,
		CreatedAt: f.CreatedAt,
	})
	if err != nil {
		return err
	}
	*f = toFaqDomain(row)
	return nil
}

func (r *FaqRepository) Get(ctx context.Context, id uuid.UUID) (*support.Faq, error) {
	row, err := r.q.GetFaq(ctx, id)
	if err != nil {
		return nil, err
	}
	f := toFaqDomain(row)
	return &f, nil
}

func (r *FaqRepository) List(ctx context.Context) ([]*support.Faq, error) {
	rows, err := r.q.ListFaqs(ctx)
	if err != nil {
		return nil, err
	}
	faqs := make([]*support.Faq, 0, len(rows))
	for _, row := range rows {
		f := toFaqDomain(row)
		faqs = append(faqs, &f)
	}
	return faqs, nil
}

func (r *FaqRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteFaq(ctx, id)
}

func toFaqDomain(row sqlc.Faq) support.Faq {
	return support.Faq{
		ID:        row.ID,
		Question:  row.Question,
		Answer:    row.Answer,
		CreatedAt: row.CreatedAt,
	}
}
