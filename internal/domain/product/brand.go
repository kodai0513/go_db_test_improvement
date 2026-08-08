package product

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Brand はブランドを表す。
type Brand struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

// BrandRepository は brands テーブルへのアクセスを抽象化する。
type BrandRepository interface {
	Create(ctx context.Context, b *Brand) error
	Get(ctx context.Context, id uuid.UUID) (*Brand, error)
	List(ctx context.Context) ([]*Brand, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
