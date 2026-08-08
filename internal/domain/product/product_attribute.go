package product

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ProductAttribute は商品属性マスタを表す。
type ProductAttribute struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

// ProductAttributeRepository は product_attributes テーブルへのアクセスを抽象化する。
type ProductAttributeRepository interface {
	Create(ctx context.Context, a *ProductAttribute) error
	Get(ctx context.Context, id uuid.UUID) (*ProductAttribute, error)
	List(ctx context.Context) ([]*ProductAttribute, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
