package loyalty

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// LoyaltyPointExpiration は会員のポイント失効予定を表す。
type LoyaltyPointExpiration struct {
	ID        uuid.UUID
	MemberID  uuid.UUID
	Points    int32
	ExpiresAt time.Time
}

// LoyaltyPointExpirationRepository は loyalty_point_expirations テーブルへのアクセスを抽象化する。
type LoyaltyPointExpirationRepository interface {
	Create(ctx context.Context, e *LoyaltyPointExpiration) error
	Get(ctx context.Context, id uuid.UUID) (*LoyaltyPointExpiration, error)
	List(ctx context.Context) ([]*LoyaltyPointExpiration, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*LoyaltyPointExpiration, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
