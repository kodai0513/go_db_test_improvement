package merchant

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// VendorPerformanceMetric はベンダーの実績指標を表す。
type VendorPerformanceMetric struct {
	ID         uuid.UUID
	VendorID   uuid.UUID
	MetricName string
	Value      int32
	RecordedAt time.Time
}

// VendorPerformanceMetricRepository は vendor_performance_metrics テーブルへのアクセスを抽象化する。
type VendorPerformanceMetricRepository interface {
	Create(ctx context.Context, m *VendorPerformanceMetric) error
	Get(ctx context.Context, id uuid.UUID) (*VendorPerformanceMetric, error)
	List(ctx context.Context) ([]*VendorPerformanceMetric, error)
	ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*VendorPerformanceMetric, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
