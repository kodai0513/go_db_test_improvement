package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres/sqlc"
)

// ProductImageRepository は product.ProductImageRepository の実装。
type ProductImageRepository struct {
	q *sqlc.Queries
}

func NewProductImageRepository(db *sql.DB) *ProductImageRepository {
	return &ProductImageRepository{q: sqlc.New(db)}
}

var _ product.ProductImageRepository = (*ProductImageRepository)(nil)

func (r *ProductImageRepository) Create(ctx context.Context, img *product.ProductImage) error {
	row, err := r.q.CreateProductImage(ctx, sqlc.CreateProductImageParams{
		ID:        img.ID,
		ProductID: img.ProductID,
		URL:       img.URL,
		SortOrder: img.SortOrder,
	})
	if err != nil {
		return err
	}
	*img = toProductImageDomain(row)
	return nil
}

func (r *ProductImageRepository) Get(ctx context.Context, id uuid.UUID) (*product.ProductImage, error) {
	row, err := r.q.GetProductImage(ctx, id)
	if err != nil {
		return nil, err
	}
	img := toProductImageDomain(row)
	return &img, nil
}

func (r *ProductImageRepository) List(ctx context.Context) ([]*product.ProductImage, error) {
	rows, err := r.q.ListProductImages(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductImage, 0, len(rows))
	for _, row := range rows {
		img := toProductImageDomain(row)
		items = append(items, &img)
	}
	return items, nil
}

func (r *ProductImageRepository) ListByProductID(ctx context.Context, productID uuid.UUID) ([]*product.ProductImage, error) {
	rows, err := r.q.ListProductImagesByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductImage, 0, len(rows))
	for _, row := range rows {
		img := toProductImageDomain(row)
		items = append(items, &img)
	}
	return items, nil
}

func (r *ProductImageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteProductImage(ctx, id)
}

func toProductImageDomain(row sqlc.ProductImage) product.ProductImage {
	return product.ProductImage{
		ID:        row.ID,
		ProductID: row.ProductID,
		URL:       row.URL,
		SortOrder: row.SortOrder,
	}
}
