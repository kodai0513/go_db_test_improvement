package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres/sqlc"
)

// SiteSettingRepository は content.SiteSettingRepository の実装。
type SiteSettingRepository struct {
	q *sqlc.Queries
}

func NewSiteSettingRepository(db *sql.DB) *SiteSettingRepository {
	return &SiteSettingRepository{q: sqlc.New(db)}
}

var _ content.SiteSettingRepository = (*SiteSettingRepository)(nil)

func (r *SiteSettingRepository) Create(ctx context.Context, s *content.SiteSetting) error {
	row, err := r.q.CreateSiteSetting(ctx, sqlc.CreateSiteSettingParams{
		ID:           s.ID,
		SettingKey:   s.SettingKey,
		SettingValue: s.SettingValue,
	})
	if err != nil {
		return err
	}
	*s = toSiteSettingDomain(row)
	return nil
}

func (r *SiteSettingRepository) Get(ctx context.Context, id uuid.UUID) (*content.SiteSetting, error) {
	row, err := r.q.GetSiteSetting(ctx, id)
	if err != nil {
		return nil, err
	}
	s := toSiteSettingDomain(row)
	return &s, nil
}

func (r *SiteSettingRepository) List(ctx context.Context) ([]*content.SiteSetting, error) {
	rows, err := r.q.ListSiteSettings(ctx)
	if err != nil {
		return nil, err
	}
	settings := make([]*content.SiteSetting, 0, len(rows))
	for _, row := range rows {
		s := toSiteSettingDomain(row)
		settings = append(settings, &s)
	}
	return settings, nil
}

func (r *SiteSettingRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteSiteSetting(ctx, id)
}

func toSiteSettingDomain(row sqlc.SiteSetting) content.SiteSetting {
	return content.SiteSetting{
		ID:           row.ID,
		SettingKey:   row.SettingKey,
		SettingValue: row.SettingValue,
	}
}
