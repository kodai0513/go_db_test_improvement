package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// Repository implements order.Repository backed by the sqlc-generated
// query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ order.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, o *order.Order) error {
	row, err := r.q.CreateOrder(ctx, sqlc.CreateOrderParams{
		ID:             o.ID,
		MemberID:       o.MemberID,
		TotalAmountYen: o.TotalAmountYen,
		Status:         o.Status,
		CreatedAt:      o.CreatedAt,
	})
	if err != nil {
		return err
	}
	*o = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*order.Order, error) {
	row, err := r.q.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}
	o := toDomain(row)
	return &o, nil
}

func (r *Repository) List(ctx context.Context) ([]*order.Order, error) {
	rows, err := r.q.ListOrders(ctx)
	if err != nil {
		return nil, err
	}
	orders := make([]*order.Order, 0, len(rows))
	for _, row := range rows {
		o := toDomain(row)
		orders = append(orders, &o)
	}
	return orders, nil
}

func (r *Repository) ListOrderSummaryByMember(ctx context.Context) ([]*order.OrderSummaryByMember, error) {
	rows, err := r.q.ListOrderSummaryByMember(ctx)
	if err != nil {
		return nil, err
	}
	summaries := make([]*order.OrderSummaryByMember, 0, len(rows))
	for _, row := range rows {
		summaries = append(summaries, &order.OrderSummaryByMember{
			MemberID:       row.MemberID,
			OrderCount:     row.OrderCount,
			TotalAmountYen: row.TotalAmountYen,
			AvgAmountYen:   row.AvgAmountYen,
		})
	}
	return summaries, nil
}

func (r *Repository) ListOrderCountByStatus(ctx context.Context) ([]*order.OrderCountByStatus, error) {
	rows, err := r.q.ListOrderCountByStatus(ctx)
	if err != nil {
		return nil, err
	}
	counts := make([]*order.OrderCountByStatus, 0, len(rows))
	for _, row := range rows {
		counts = append(counts, &order.OrderCountByStatus{
			Status:     row.Status,
			OrderCount: row.OrderCount,
		})
	}
	return counts, nil
}

func (r *Repository) ListSalesQuantityAndAmountByProduct(ctx context.Context) ([]*order.SalesQuantityAndAmountByProduct, error) {
	rows, err := r.q.ListSalesQuantityAndAmountByProduct(ctx)
	if err != nil {
		return nil, err
	}
	sales := make([]*order.SalesQuantityAndAmountByProduct, 0, len(rows))
	for _, row := range rows {
		sales = append(sales, &order.SalesQuantityAndAmountByProduct{
			ProductID:      row.ProductID,
			TotalQuantity:  row.TotalQuantity,
			TotalAmountYen: row.TotalAmountYen,
		})
	}
	return sales, nil
}

func toDomain(row sqlc.Order) order.Order {
	return order.Order{
		ID:             row.ID,
		MemberID:       row.MemberID,
		TotalAmountYen: row.TotalAmountYen,
		Status:         row.Status,
		CreatedAt:      row.CreatedAt,
	}
}
