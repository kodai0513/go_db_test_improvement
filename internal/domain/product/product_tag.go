package product

import (
	"context"

	"github.com/google/uuid"
)

// ProductTag は商品タグマスタを表す。
type ProductTag struct {
	ID   uuid.UUID
	Name string
}

// ProductTagRepository は product_tags テーブルへのアクセスを抽象化する。
type ProductTagRepository interface {
	Create(ctx context.Context, t *ProductTag) error
	Get(ctx context.Context, id uuid.UUID) (*ProductTag, error)
	List(ctx context.Context) ([]*ProductTag, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
