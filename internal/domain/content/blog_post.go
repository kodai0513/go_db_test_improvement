package content

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// BlogPost はブログ記事を表す。
// published_at はDB上はNULL許容だが、Go側は未公開の場合ゼロ値のtime.Timeを書き込む運用とする。
type BlogPost struct {
	ID          uuid.UUID
	Title       string
	Body        string
	PublishedAt time.Time
}

// BlogPostRepository は blog_posts テーブルへのアクセスを抽象化する。
type BlogPostRepository interface {
	Create(ctx context.Context, p *BlogPost) error
	Get(ctx context.Context, id uuid.UUID) (*BlogPost, error)
	List(ctx context.Context) ([]*BlogPost, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
