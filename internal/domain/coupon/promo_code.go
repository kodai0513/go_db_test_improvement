package coupon

import (
	"context"

	"github.com/google/uuid"
)

// PromoCode はキャンペーンに紐づくプロモーションコードを表す。
type PromoCode struct {
	ID         uuid.UUID
	Code       string
	CampaignID uuid.UUID
}

// PromoCodeRepository は promo_codes テーブルへのアクセスを抽象化する。
type PromoCodeRepository interface {
	Create(ctx context.Context, p *PromoCode) error
	Get(ctx context.Context, id uuid.UUID) (*PromoCode, error)
	List(ctx context.Context) ([]*PromoCode, error)
	ListByCampaignID(ctx context.Context, campaignID uuid.UUID) ([]*PromoCode, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
