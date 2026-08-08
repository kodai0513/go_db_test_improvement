package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres/sqlc"
)

// BlogPostRepository は content.BlogPostRepository の実装。
type BlogPostRepository struct {
	q *sqlc.Queries
}

func NewBlogPostRepository(db *sql.DB) *BlogPostRepository {
	return &BlogPostRepository{q: sqlc.New(db)}
}

var _ content.BlogPostRepository = (*BlogPostRepository)(nil)

func (r *BlogPostRepository) Create(ctx context.Context, p *content.BlogPost) error {
	row, err := r.q.CreateBlogPost(ctx, sqlc.CreateBlogPostParams{
		ID:          p.ID,
		Title:       p.Title,
		Body:        p.Body,
		PublishedAt: p.PublishedAt,
	})
	if err != nil {
		return err
	}
	*p = toBlogPostDomain(row)
	return nil
}

func (r *BlogPostRepository) Get(ctx context.Context, id uuid.UUID) (*content.BlogPost, error) {
	row, err := r.q.GetBlogPost(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toBlogPostDomain(row)
	return &p, nil
}

func (r *BlogPostRepository) List(ctx context.Context) ([]*content.BlogPost, error) {
	rows, err := r.q.ListBlogPosts(ctx)
	if err != nil {
		return nil, err
	}
	posts := make([]*content.BlogPost, 0, len(rows))
	for _, row := range rows {
		p := toBlogPostDomain(row)
		posts = append(posts, &p)
	}
	return posts, nil
}

func (r *BlogPostRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteBlogPost(ctx, id)
}

func toBlogPostDomain(row sqlc.BlogPost) content.BlogPost {
	return content.BlogPost{
		ID:          row.ID,
		Title:       row.Title,
		Body:        row.Body,
		PublishedAt: row.PublishedAt,
	}
}
