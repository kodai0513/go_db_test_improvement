package loyalty

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// PointRedemption は会員によるポイント交換を表す。
type PointRedemption struct {
	ID         uuid.UUID
	MemberID   uuid.UUID
	Points     int32
	RedeemedAt time.Time
}

// PointRedemptionRepository は point_redemptions テーブルへのアクセスを抽象化する。
type PointRedemptionRepository interface {
	Create(ctx context.Context, r *PointRedemption) error
	Get(ctx context.Context, id uuid.UUID) (*PointRedemption, error)
	List(ctx context.Context) ([]*PointRedemption, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*PointRedemption, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
