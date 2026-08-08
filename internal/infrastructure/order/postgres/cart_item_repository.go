package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// CartItemRepository は order.CartItemRepository の実装。
type CartItemRepository struct {
	q *sqlc.Queries
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{q: sqlc.New(db)}
}

var _ order.CartItemRepository = (*CartItemRepository)(nil)

func (r *CartItemRepository) Create(ctx context.Context, i *order.CartItem) error {
	row, err := r.q.CreateCartItem(ctx, sqlc.CreateCartItemParams{
		ID:        i.ID,
		CartID:    i.CartID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
	})
	if err != nil {
		return err
	}
	*i = toCartItemDomain(row)
	return nil
}

func (r *CartItemRepository) Get(ctx context.Context, id uuid.UUID) (*order.CartItem, error) {
	row, err := r.q.GetCartItem(ctx, id)
	if err != nil {
		return nil, err
	}
	i := toCartItemDomain(row)
	return &i, nil
}

func (r *CartItemRepository) List(ctx context.Context) ([]*order.CartItem, error) {
	rows, err := r.q.ListCartItems(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*order.CartItem, 0, len(rows))
	for _, row := range rows {
		i := toCartItemDomain(row)
		items = append(items, &i)
	}
	return items, nil
}

func (r *CartItemRepository) ListByCartID(ctx context.Context, cartID uuid.UUID) ([]*order.CartItem, error) {
	rows, err := r.q.ListCartItemsByCartID(ctx, cartID)
	if err != nil {
		return nil, err
	}
	items := make([]*order.CartItem, 0, len(rows))
	for _, row := range rows {
		i := toCartItemDomain(row)
		items = append(items, &i)
	}
	return items, nil
}

func (r *CartItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCartItem(ctx, id)
}

func toCartItemDomain(row sqlc.CartItem) order.CartItem {
	return order.CartItem{
		ID:        row.ID,
		CartID:    row.CartID,
		ProductID: row.ProductID,
		Quantity:  row.Quantity,
	}
}
