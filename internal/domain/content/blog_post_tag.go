package content

import (
	"context"

	"github.com/google/uuid"
)

// BlogPostTag はブログ記事に付与されたタグを表す。
type BlogPostTag struct {
	ID         uuid.UUID
	BlogPostID uuid.UUID
	Tag        string
}

// BlogPostTagRepository は blog_post_tags テーブルへのアクセスを抽象化する。
type BlogPostTagRepository interface {
	Create(ctx context.Context, t *BlogPostTag) error
	Get(ctx context.Context, id uuid.UUID) (*BlogPostTag, error)
	List(ctx context.Context) ([]*BlogPostTag, error)
	ListByBlogPostID(ctx context.Context, blogPostID uuid.UUID) ([]*BlogPostTag, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
