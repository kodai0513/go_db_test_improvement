package loyalty

import (
	"context"

	"github.com/google/uuid"
)

// RewardItem は交換可能な報酬アイテムを表す。
type RewardItem struct {
	ID         uuid.UUID
	Name       string
	PointsCost int32
}

// RewardItemRepository は reward_items テーブルへのアクセスを抽象化する。
type RewardItemRepository interface {
	Create(ctx context.Context, i *RewardItem) error
	Get(ctx context.Context, id uuid.UUID) (*RewardItem, error)
	List(ctx context.Context) ([]*RewardItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
