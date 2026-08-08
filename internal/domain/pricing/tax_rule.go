package pricing

import (
	"context"

	"github.com/google/uuid"
)

// TaxRule は税率が適用される地域を表す。
type TaxRule struct {
	ID        uuid.UUID
	TaxRateID uuid.UUID
	Region    string
}

// TaxRuleRepository は tax_rules テーブルへのアクセスを抽象化する。
type TaxRuleRepository interface {
	Create(ctx context.Context, t *TaxRule) error
	Get(ctx context.Context, id uuid.UUID) (*TaxRule, error)
	List(ctx context.Context) ([]*TaxRule, error)
	ListByTaxRateID(ctx context.Context, taxRateID uuid.UUID) ([]*TaxRule, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
