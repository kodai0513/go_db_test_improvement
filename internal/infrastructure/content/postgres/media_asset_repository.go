package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres/sqlc"
)

// MediaAssetRepository は content.MediaAssetRepository の実装。
type MediaAssetRepository struct {
	q *sqlc.Queries
}

func NewMediaAssetRepository(db *sql.DB) *MediaAssetRepository {
	return &MediaAssetRepository{q: sqlc.New(db)}
}

var _ content.MediaAssetRepository = (*MediaAssetRepository)(nil)

func (r *MediaAssetRepository) Create(ctx context.Context, m *content.MediaAsset) error {
	row, err := r.q.CreateMediaAsset(ctx, sqlc.CreateMediaAssetParams{
		ID:        m.ID,
		Url:       m.URL,
		MediaType: m.MediaType,
		CreatedAt: m.CreatedAt,
	})
	if err != nil {
		return err
	}
	*m = toMediaAssetDomain(row)
	return nil
}

func (r *MediaAssetRepository) Get(ctx context.Context, id uuid.UUID) (*content.MediaAsset, error) {
	row, err := r.q.GetMediaAsset(ctx, id)
	if err != nil {
		return nil, err
	}
	m := toMediaAssetDomain(row)
	return &m, nil
}

func (r *MediaAssetRepository) List(ctx context.Context) ([]*content.MediaAsset, error) {
	rows, err := r.q.ListMediaAssets(ctx)
	if err != nil {
		return nil, err
	}
	assets := make([]*content.MediaAsset, 0, len(rows))
	for _, row := range rows {
		m := toMediaAssetDomain(row)
		assets = append(assets, &m)
	}
	return assets, nil
}

func (r *MediaAssetRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMediaAsset(ctx, id)
}

func toMediaAssetDomain(row sqlc.MediaAsset) content.MediaAsset {
	return content.MediaAsset{
		ID:        row.ID,
		URL:       row.Url,
		MediaType: row.MediaType,
		CreatedAt: row.CreatedAt,
	}
}
