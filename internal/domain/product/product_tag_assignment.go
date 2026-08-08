package product

import (
	"context"

	"github.com/google/uuid"
)

// ProductTagAssignment は商品とタグの割り当てを表す。
type ProductTagAssignment struct {
	ID           uuid.UUID
	ProductID    uuid.UUID
	ProductTagID uuid.UUID
}

// ProductTagAssignmentRepository は product_tag_assignments テーブルへのアクセスを抽象化する。
type ProductTagAssignmentRepository interface {
	Create(ctx context.Context, a *ProductTagAssignment) error
	Get(ctx context.Context, id uuid.UUID) (*ProductTagAssignment, error)
	List(ctx context.Context) ([]*ProductTagAssignment, error)
	ListByProductID(ctx context.Context, productID uuid.UUID) ([]*ProductTagAssignment, error)
	ListByProductTagID(ctx context.Context, productTagID uuid.UUID) ([]*ProductTagAssignment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
