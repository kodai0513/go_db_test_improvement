package coupon

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Campaign はクーポンを配布するキャンペーンを表す。
type Campaign struct {
	ID       uuid.UUID
	Name     string
	StartsAt time.Time
	EndsAt   time.Time
}

// CampaignRepository は campaigns テーブルへのアクセスを抽象化する。
type CampaignRepository interface {
	Create(ctx context.Context, c *Campaign) error
	Get(ctx context.Context, id uuid.UUID) (*Campaign, error)
	List(ctx context.Context) ([]*Campaign, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
