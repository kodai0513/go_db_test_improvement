package loyalty

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// LoyaltyProgram はロイヤリティプログラムを表す。
type LoyaltyProgram struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

// LoyaltyProgramRepository は loyalty_programs テーブルへのアクセスを抽象化する。
type LoyaltyProgramRepository interface {
	Create(ctx context.Context, p *LoyaltyProgram) error
	Get(ctx context.Context, id uuid.UUID) (*LoyaltyProgram, error)
	List(ctx context.Context) ([]*LoyaltyProgram, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
