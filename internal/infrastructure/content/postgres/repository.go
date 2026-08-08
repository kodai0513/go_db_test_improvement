package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres/sqlc"
)

// Repository は content.BannerRepository の実装。
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ content.BannerRepository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, b *content.Banner) error {
	row, err := r.q.CreateBanner(ctx, sqlc.CreateBannerParams{
		ID:       b.ID,
		Title:    b.Title,
		ImageUrl: b.ImageURL,
		StartsAt: b.StartsAt,
		EndsAt:   b.EndsAt,
	})
	if err != nil {
		return err
	}
	*b = toBannerDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*content.Banner, error) {
	row, err := r.q.GetBanner(ctx, id)
	if err != nil {
		return nil, err
	}
	b := toBannerDomain(row)
	return &b, nil
}

func (r *Repository) List(ctx context.Context) ([]*content.Banner, error) {
	rows, err := r.q.ListBanners(ctx)
	if err != nil {
		return nil, err
	}
	banners := make([]*content.Banner, 0, len(rows))
	for _, row := range rows {
		b := toBannerDomain(row)
		banners = append(banners, &b)
	}
	return banners, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteBanner(ctx, id)
}

func toBannerDomain(row sqlc.Banner) content.Banner {
	return content.Banner{
		ID:       row.ID,
		Title:    row.Title,
		ImageURL: row.ImageUrl,
		StartsAt: row.StartsAt,
		EndsAt:   row.EndsAt,
	}
}
