package product

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, p *Product) error
	Get(ctx context.Context, id uuid.UUID) (*Product, error)
	List(ctx context.Context) ([]*Product, error)

	// ListProductCountAndAvgPriceByCategory はカテゴリごとの商品数と平均価格を集計する。
	ListProductCountAndAvgPriceByCategory(ctx context.Context) ([]*ProductCategorySummary, error)
	// ListVariantPriceRangeByProduct は商品ごとのバリエーション数と価格レンジ(最安値・最高値)を集計する。
	ListVariantPriceRangeByProduct(ctx context.Context) ([]*ProductVariantPriceRange, error)
	// ListProductCountByTag は商品タグごとの商品数を集計する。
	ListProductCountByTag(ctx context.Context) ([]*ProductTagSummary, error)
}

// ProductCategorySummary はカテゴリごとの商品数と平均価格の集計結果を表す。
type ProductCategorySummary struct {
	CategoryID   uuid.UUID
	CategoryName string
	ProductCount int64
	AvgPriceYen  float64
}

// ProductVariantPriceRange は商品ごとのバリエーション数と価格レンジ(最安値・最高値)の集計結果を表す。
type ProductVariantPriceRange struct {
	ProductID    uuid.UUID
	ProductName  string
	VariantCount int64
	MinPriceYen  int32
	MaxPriceYen  int32
}

// ProductTagSummary は商品タグごとの商品数の集計結果を表す。
type ProductTagSummary struct {
	ProductTagID uuid.UUID
	TagName      string
	ProductCount int64
}
