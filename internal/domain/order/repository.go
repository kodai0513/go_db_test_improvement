package order

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, o *Order) error
	Get(ctx context.Context, id uuid.UUID) (*Order, error)
	List(ctx context.Context) ([]*Order, error)
	ListOrderSummaryByMember(ctx context.Context) ([]*OrderSummaryByMember, error)
	ListOrderCountByStatus(ctx context.Context) ([]*OrderCountByStatus, error)
	ListSalesQuantityAndAmountByProduct(ctx context.Context) ([]*SalesQuantityAndAmountByProduct, error)
}

// OrderSummaryByMember は会員ごとの注文件数・合計金額・平均注文金額の集計結果を表す。
type OrderSummaryByMember struct {
	MemberID       uuid.UUID
	OrderCount     int64
	TotalAmountYen int64
	AvgAmountYen   float64
}

// OrderCountByStatus はステータスごとの注文件数の集計結果を表す。
type OrderCountByStatus struct {
	Status     string
	OrderCount int64
}

// SalesQuantityAndAmountByProduct は商品ごとの販売数量・売上合計の集計結果を表す。
type SalesQuantityAndAmountByProduct struct {
	ProductID      uuid.UUID
	TotalQuantity  int64
	TotalAmountYen int64
}
