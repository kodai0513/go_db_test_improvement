package product

import (
	"context"

	"github.com/google/uuid"
)

// ProductAttributeValue は商品ごとの属性値を表す。
type ProductAttributeValue struct {
	ID                  uuid.UUID
	ProductID           uuid.UUID
	ProductAttributeID  uuid.UUID
	Value               string
}

// ProductAttributeValueRepository は product_attribute_values テーブルへのアクセスを抽象化する。
type ProductAttributeValueRepository interface {
	Create(ctx context.Context, v *ProductAttributeValue) error
	Get(ctx context.Context, id uuid.UUID) (*ProductAttributeValue, error)
	List(ctx context.Context) ([]*ProductAttributeValue, error)
	ListByProductID(ctx context.Context, productID uuid.UUID) ([]*ProductAttributeValue, error)
	ListByProductAttributeID(ctx context.Context, productAttributeID uuid.UUID) ([]*ProductAttributeValue, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
