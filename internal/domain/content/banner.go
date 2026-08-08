package content

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Banner はトップページ等に表示するバナーを表す。
type Banner struct {
	ID       uuid.UUID
	Title    string
	ImageURL string
	StartsAt time.Time
	EndsAt   time.Time
}

// BannerRepository は banners テーブルへのアクセスを抽象化する。
type BannerRepository interface {
	Create(ctx context.Context, b *Banner) error
	Get(ctx context.Context, id uuid.UUID) (*Banner, error)
	List(ctx context.Context) ([]*Banner, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
