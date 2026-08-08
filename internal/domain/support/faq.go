package support

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Faq はよくある質問(FAQ)を表す。
type Faq struct {
	ID        uuid.UUID
	Question  string
	Answer    string
	CreatedAt time.Time
}

// FaqRepository は faqs テーブルへのアクセスを抽象化する。
type FaqRepository interface {
	Create(ctx context.Context, f *Faq) error
	Get(ctx context.Context, id uuid.UUID) (*Faq, error)
	List(ctx context.Context) ([]*Faq, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
