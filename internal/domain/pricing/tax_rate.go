package pricing

import (
	"context"

	"github.com/google/uuid"
)

// TaxRate は税率を表す。
type TaxRate struct {
	ID   uuid.UUID
	Name string
	Rate int32
}

// TaxRateRepository は tax_rates テーブルへのアクセスを抽象化する。
type TaxRateRepository interface {
	Create(ctx context.Context, t *TaxRate) error
	Get(ctx context.Context, id uuid.UUID) (*TaxRate, error)
	List(ctx context.Context) ([]*TaxRate, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
