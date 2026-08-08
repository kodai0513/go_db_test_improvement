package merchant

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Vendor は出品事業者を表す。
type Vendor struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
}

// VendorRepository は vendors テーブルへのアクセスを抽象化する。
type VendorRepository interface {
	Create(ctx context.Context, v *Vendor) error
	Get(ctx context.Context, id uuid.UUID) (*Vendor, error)
	List(ctx context.Context) ([]*Vendor, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
