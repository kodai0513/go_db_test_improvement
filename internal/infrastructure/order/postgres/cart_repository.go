package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// CartRepository は order.CartRepository の実装。
type CartRepository struct {
	q *sqlc.Queries
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{q: sqlc.New(db)}
}

var _ order.CartRepository = (*CartRepository)(nil)

func (r *CartRepository) Create(ctx context.Context, c *order.Cart) error {
	row, err := r.q.CreateCart(ctx, sqlc.CreateCartParams{
		ID:        c.ID,
		MemberID:  c.MemberID,
		CreatedAt: c.CreatedAt,
	})
	if err != nil {
		return err
	}
	*c = toCartDomain(row)
	return nil
}

func (r *CartRepository) Get(ctx context.Context, id uuid.UUID) (*order.Cart, error) {
	row, err := r.q.GetCart(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toCartDomain(row)
	return &c, nil
}

func (r *CartRepository) List(ctx context.Context) ([]*order.Cart, error) {
	rows, err := r.q.ListCarts(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*order.Cart, 0, len(rows))
	for _, row := range rows {
		c := toCartDomain(row)
		items = append(items, &c)
	}
	return items, nil
}

func (r *CartRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*order.Cart, error) {
	rows, err := r.q.ListCartsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	items := make([]*order.Cart, 0, len(rows))
	for _, row := range rows {
		c := toCartDomain(row)
		items = append(items, &c)
	}
	return items, nil
}

func (r *CartRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCart(ctx, id)
}

func toCartDomain(row sqlc.Cart) order.Cart {
	return order.Cart{
		ID:        row.ID,
		MemberID:  row.MemberID,
		CreatedAt: row.CreatedAt,
	}
}
