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

func toDomain(row sqlc.Product) product.Product {
	return product.Product{
		ID:        row.ID,
		Name:      row.Name,
		PriceYen:  row.PriceYen,
		CreatedAt: row.CreatedAt,
	}
}
