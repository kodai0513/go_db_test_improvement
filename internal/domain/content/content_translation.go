package content

import (
	"context"

	"github.com/google/uuid"
)

// ContentTranslation はページ本文の多言語翻訳を表す。
type ContentTranslation struct {
	ID             uuid.UUID
	PageID         uuid.UUID
	Locale         string
	TranslatedBody string
}

// ContentTranslationRepository は content_translations テーブルへのアクセスを抽象化する。
type ContentTranslationRepository interface {
	Create(ctx context.Context, t *ContentTranslation) error
	Get(ctx context.Context, id uuid.UUID) (*ContentTranslation, error)
	List(ctx context.Context) ([]*ContentTranslation, error)
	ListByPageID(ctx context.Context, pageID uuid.UUID) ([]*ContentTranslation, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
