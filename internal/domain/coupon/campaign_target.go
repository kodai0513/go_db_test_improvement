package coupon

import (
	"context"

	"github.com/google/uuid"
)

// CampaignTarget はキャンペーンの対象会員を表す。
type CampaignTarget struct {
	ID         uuid.UUID
	CampaignID uuid.UUID
	MemberID   uuid.UUID
}

// CampaignTargetRepository は campaign_targets テーブルへのアクセスを抽象化する。
type CampaignTargetRepository interface {
	Create(ctx context.Context, t *CampaignTarget) error
	Get(ctx context.Context, id uuid.UUID) (*CampaignTarget, error)
	List(ctx context.Context) ([]*CampaignTarget, error)
	ListByCampaignID(ctx context.Context, campaignID uuid.UUID) ([]*CampaignTarget, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
