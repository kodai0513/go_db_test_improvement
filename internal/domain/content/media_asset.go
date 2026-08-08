package content

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MediaAsset は画像や動画などのメディアファイルを表す。
type MediaAsset struct {
	ID        uuid.UUID
	URL       string
	MediaType string
	CreatedAt time.Time
}

// MediaAssetRepository は media_assets テーブルへのアクセスを抽象化する。
type MediaAssetRepository interface {
	Create(ctx context.Context, m *MediaAsset) error
	Get(ctx context.Context, id uuid.UUID) (*MediaAsset, error)
	List(ctx context.Context) ([]*MediaAsset, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
