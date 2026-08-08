package member

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MemberTag は会員に付与できるタグを表す。
type MemberTag struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

// MemberTagRepository は member_tags テーブルへのアクセスを抽象化する。
type MemberTagRepository interface {
	Create(ctx context.Context, t *MemberTag) error
	Get(ctx context.Context, id uuid.UUID) (*MemberTag, error)
	List(ctx context.Context) ([]*MemberTag, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
