package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres/sqlc"
)

// BlogPostTagRepository は content.BlogPostTagRepository の実装。
type BlogPostTagRepository struct {
	q *sqlc.Queries
}

func NewBlogPostTagRepository(db *sql.DB) *BlogPostTagRepository {
	return &BlogPostTagRepository{q: sqlc.New(db)}
}

var _ content.BlogPostTagRepository = (*BlogPostTagRepository)(nil)

func (r *BlogPostTagRepository) Create(ctx context.Context, t *content.BlogPostTag) error {
	row, err := r.q.CreateBlogPostTag(ctx, sqlc.CreateBlogPostTagParams{
		ID:         t.ID,
		BlogPostID: t.BlogPostID,
		Tag:        t.Tag,
	})
	if err != nil {
		return err
	}
	*t = toBlogPostTagDomain(row)
	return nil
}

func (r *BlogPostTagRepository) Get(ctx context.Context, id uuid.UUID) (*content.BlogPostTag, error) {
	row, err := r.q.GetBlogPostTag(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toBlogPostTagDomain(row)
	return &t, nil
}

func (r *BlogPostTagRepository) List(ctx context.Context) ([]*content.BlogPostTag, error) {
	rows, err := r.q.ListBlogPostTags(ctx)
	if err != nil {
		return nil, err
	}
	tags := make([]*content.BlogPostTag, 0, len(rows))
	for _, row := range rows {
		t := toBlogPostTagDomain(row)
		tags = append(tags, &t)
	}
	return tags, nil
}

func (r *BlogPostTagRepository) ListByBlogPostID(ctx context.Context, blogPostID uuid.UUID) ([]*content.BlogPostTag, error) {
	rows, err := r.q.ListBlogPostTagsByBlogPostID(ctx, blogPostID)
	if err != nil {
		return nil, err
	}
	tags := make([]*content.BlogPostTag, 0, len(rows))
	for _, row := range rows {
		t := toBlogPostTagDomain(row)
		tags = append(tags, &t)
	}
	return tags, nil
}

func (r *BlogPostTagRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteBlogPostTag(ctx, id)
}

func toBlogPostTagDomain(row sqlc.BlogPostTag) content.BlogPostTag {
	return content.BlogPostTag{
		ID:         row.ID,
		BlogPostID: row.BlogPostID,
		Tag:        row.Tag,
	}
}
