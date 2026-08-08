package loyalty

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// BirthdayReward は誕生日特典によるポイント付与を表す。
type BirthdayReward struct {
	ID        uuid.UUID
	MemberID  uuid.UUID
	Points    int32
	GrantedAt time.Time
}

// BirthdayRewardRepository は birthday_rewards テーブルへのアクセスを抽象化する。
type BirthdayRewardRepository interface {
	Create(ctx context.Context, b *BirthdayReward) error
	Get(ctx context.Context, id uuid.UUID) (*BirthdayReward, error)
	List(ctx context.Context) ([]*BirthdayReward, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*BirthdayReward, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
