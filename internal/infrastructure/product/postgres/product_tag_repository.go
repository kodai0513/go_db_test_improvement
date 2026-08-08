package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres/sqlc"
)

// ProductTagRepository は product.ProductTagRepository の実装。
type ProductTagRepository struct {
	q *sqlc.Queries
}

func NewProductTagRepository(db *sql.DB) *ProductTagRepository {
	return &ProductTagRepository{q: sqlc.New(db)}
}

var _ product.ProductTagRepository = (*ProductTagRepository)(nil)

func (r *ProductTagRepository) Create(ctx context.Context, t *product.ProductTag) error {
	row, err := r.q.CreateProductTag(ctx, sqlc.CreateProductTagParams{
		ID:   t.ID,
		Name: t.Name,
	})
	if err != nil {
		return err
	}
	*t = toProductTagDomain(row)
	return nil
}

func (r *ProductTagRepository) Get(ctx context.Context, id uuid.UUID) (*product.ProductTag, error) {
	row, err := r.q.GetProductTag(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toProductTagDomain(row)
	return &t, nil
}

func (r *ProductTagRepository) List(ctx context.Context) ([]*product.ProductTag, error) {
	rows, err := r.q.ListProductTags(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductTag, 0, len(rows))
	for _, row := range rows {
		t := toProductTagDomain(row)
		items = append(items, &t)
	}
	return items, nil
}

func (r *ProductTagRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteProductTag(ctx, id)
}

func toProductTagDomain(row sqlc.ProductTag) product.ProductTag {
	return product.ProductTag{
		ID:   row.ID,
		Name: row.Name,
	}
}
