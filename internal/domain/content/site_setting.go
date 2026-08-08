package content

import (
	"context"

	"github.com/google/uuid"
)

// SiteSetting はサイト全体の設定値を表す。
type SiteSetting struct {
	ID           uuid.UUID
	SettingKey   string
	SettingValue string
}

// SiteSettingRepository は site_settings テーブルへのアクセスを抽象化する。
type SiteSettingRepository interface {
	Create(ctx context.Context, s *SiteSetting) error
	Get(ctx context.Context, id uuid.UUID) (*SiteSetting, error)
	List(ctx context.Context) ([]*SiteSetting, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
