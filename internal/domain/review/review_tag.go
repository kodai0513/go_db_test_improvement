package review

import (
	"context"

	"github.com/google/uuid"
)

// ReviewTag はレビューに付与できるタグを表す。
type ReviewTag struct {
	ID   uuid.UUID
	Name string
}

// ReviewTagRepository は review_tags テーブルへのアクセスを抽象化する。
type ReviewTagRepository interface {
	Create(ctx context.Context, t *ReviewTag) error
	Get(ctx context.Context, id uuid.UUID) (*ReviewTag, error)
	List(ctx context.Context) ([]*ReviewTag, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
