package pricing

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// CurrencyRate は通貨ごとの対円レートを表す。
type CurrencyRate struct {
	ID          uuid.UUID
	CurrencyID  uuid.UUID
	RateToJPY   int32
	EffectiveAt time.Time
}

// CurrencyRateRepository は currency_rates テーブルへのアクセスを抽象化する。
type CurrencyRateRepository interface {
	Create(ctx context.Context, c *CurrencyRate) error
	Get(ctx context.Context, id uuid.UUID) (*CurrencyRate, error)
	List(ctx context.Context) ([]*CurrencyRate, error)
	ListByCurrencyID(ctx context.Context, currencyID uuid.UUID) ([]*CurrencyRate, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
