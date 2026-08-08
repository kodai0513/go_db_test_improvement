package inventory

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Supplier は仕入先を表す。
type Supplier struct {
	ID           uuid.UUID
	Name         string
	ContactEmail string
	CreatedAt    time.Time
}

// SupplierRepository は suppliers テーブルへのアクセスを抽象化する。
type SupplierRepository interface {
	Create(ctx context.Context, s *Supplier) error
	Get(ctx context.Context, id uuid.UUID) (*Supplier, error)
	List(ctx context.Context) ([]*Supplier, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
