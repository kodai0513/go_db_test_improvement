package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres/sqlc"
)

// ProductTagAssignmentRepository は product.ProductTagAssignmentRepository の実装。
type ProductTagAssignmentRepository struct {
	q *sqlc.Queries
}

func NewProductTagAssignmentRepository(db *sql.DB) *ProductTagAssignmentRepository {
	return &ProductTagAssignmentRepository{q: sqlc.New(db)}
}

var _ product.ProductTagAssignmentRepository = (*ProductTagAssignmentRepository)(nil)

func (r *ProductTagAssignmentRepository) Create(ctx context.Context, a *product.ProductTagAssignment) error {
	row, err := r.q.CreateProductTagAssignment(ctx, sqlc.CreateProductTagAssignmentParams{
		ID:           a.ID,
		ProductID:    a.ProductID,
		ProductTagID: a.ProductTagID,
	})
	if err != nil {
		return err
	}
	*a = toProductTagAssignmentDomain(row)
	return nil
}

func (r *ProductTagAssignmentRepository) Get(ctx context.Context, id uuid.UUID) (*product.ProductTagAssignment, error) {
	row, err := r.q.GetProductTagAssignment(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toProductTagAssignmentDomain(row)
	return &a, nil
}

func (r *ProductTagAssignmentRepository) List(ctx context.Context) ([]*product.ProductTagAssignment, error) {
	rows, err := r.q.ListProductTagAssignments(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductTagAssignment, 0, len(rows))
	for _, row := range rows {
		a := toProductTagAssignmentDomain(row)
		items = append(items, &a)
	}
	return items, nil
}

func (r *ProductTagAssignmentRepository) ListByProductID(ctx context.Context, productID uuid.UUID) ([]*product.ProductTagAssignment, error) {
	rows, err := r.q.ListProductTagAssignmentsByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductTagAssignment, 0, len(rows))
	for _, row := range rows {
		a := toProductTagAssignmentDomain(row)
		items = append(items, &a)
	}
	return items, nil
}

func (r *ProductTagAssignmentRepository) ListByProductTagID(ctx context.Context, productTagID uuid.UUID) ([]*product.ProductTagAssignment, error) {
	rows, err := r.q.ListProductTagAssignmentsByProductTagID(ctx, productTagID)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductTagAssignment, 0, len(rows))
	for _, row := range rows {
		a := toProductTagAssignmentDomain(row)
		items = append(items, &a)
	}
	return items, nil
}

func (r *ProductTagAssignmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteProductTagAssignment(ctx, id)
}

func toProductTagAssignmentDomain(row sqlc.ProductTagAssignment) product.ProductTagAssignment {
	return product.ProductTagAssignment{
		ID:           row.ID,
		ProductID:    row.ProductID,
		ProductTagID: row.ProductTagID,
	}
}
