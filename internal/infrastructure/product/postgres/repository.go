package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres/sqlc"
)

// Repository implements product.Repository backed by the sqlc-generated
// query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ product.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, p *product.Product) error {
	row, err := r.q.CreateProduct(ctx, sqlc.CreateProductParams{
		ID:        p.ID,
		Name:      p.Name,
		PriceYen:  p.PriceYen,
		CreatedAt: p.CreatedAt,
	})
	if err != nil {
		return err
	}
	*p = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*product.Product, error) {
	row, err := r.q.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toDomain(row)
	return &p, nil
}

func (r *Repository) List(ctx context.Context) ([]*product.Product, error) {
	rows, err := r.q.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	products := make([]*product.Product, 0, len(rows))
	for _, row := range rows {
		p := toDomain(row)
		products = append(products, &p)
	}
	return products, nil
}

func (r *Repository) ListProductCountAndAvgPriceByCategory(ctx context.Context) ([]*product.ProductCategorySummary, error) {
	rows, err := r.q.ListProductCountAndAvgPriceByCategory(ctx)
	if err != nil {
		return nil, err
	}
	summaries := make([]*product.ProductCategorySummary, 0, len(rows))
	for _, row := range rows {
		summaries = append(summaries, &product.ProductCategorySummary{
			CategoryID:   row.CategoryID,
			CategoryName: row.CategoryName,
			ProductCount: row.ProductCount,
			AvgPriceYen:  row.AvgPriceYen,
		})
	}
	return summaries, nil
}

func (r *Repository) ListVariantPriceRangeByProduct(ctx context.Context) ([]*product.ProductVariantPriceRange, error) {
	rows, err := r.q.ListVariantPriceRangeByProduct(ctx)
	if err != nil {
		return nil, err
	}
	ranges := make([]*product.ProductVariantPriceRange, 0, len(rows))
	for _, row := range rows {
		ranges = append(ranges, &product.ProductVariantPriceRange{
			ProductID:    row.ProductID,
			ProductName:  row.ProductName,
			VariantCount: row.VariantCount,
			MinPriceYen:  row.MinPriceYen,
			MaxPriceYen:  row.MaxPriceYen,
		})
	}
	return ranges, nil
}

func (r *Repository) ListProductCountByTag(ctx context.Context) ([]*product.ProductTagSummary, error) {
	rows, err := r.q.ListProductCountByTag(ctx)
	if err != nil {
		return nil, err
	}
	summaries := make([]*product.ProductTagSummary, 0, len(rows))
	for _, row := range rows {
		summaries = append(summaries, &product.ProductTagSummary{
			ProductTagID: row.ProductTagID,
			TagName:      row.TagName,
			ProductCount: row.ProductCount,
		})
	}
	return summaries, nil
}

func toDomain(row sqlc.Product) product.Product {
	return product.Product{
		ID:        row.ID,
		Name:      row.Name,
		PriceYen:  row.PriceYen,
		CreatedAt: row.CreatedAt,
	}
}
