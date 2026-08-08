package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres/sqlc"
)

// VendorPerformanceMetricRepository は merchant.VendorPerformanceMetricRepository の実装。
type VendorPerformanceMetricRepository struct {
	q *sqlc.Queries
}

func NewVendorPerformanceMetricRepository(db *sql.DB) *VendorPerformanceMetricRepository {
	return &VendorPerformanceMetricRepository{q: sqlc.New(db)}
}

var _ merchant.VendorPerformanceMetricRepository = (*VendorPerformanceMetricRepository)(nil)

func (r *VendorPerformanceMetricRepository) Create(ctx context.Context, m *merchant.VendorPerformanceMetric) error {
	row, err := r.q.CreateVendorPerformanceMetric(ctx, sqlc.CreateVendorPerformanceMetricParams{
		ID:         m.ID,
		VendorID:   m.VendorID,
		MetricName: m.MetricName,
		Value:      m.Value,
		RecordedAt: m.RecordedAt,
	})
	if err != nil {
		return err
	}
	*m = toVendorPerformanceMetricDomain(row)
	return nil
}

func (r *VendorPerformanceMetricRepository) Get(ctx context.Context, id uuid.UUID) (*merchant.VendorPerformanceMetric, error) {
	row, err := r.q.GetVendorPerformanceMetric(ctx, id)
	if err != nil {
		return nil, err
	}
	m := toVendorPerformanceMetricDomain(row)
	return &m, nil
}

func (r *VendorPerformanceMetricRepository) List(ctx context.Context) ([]*merchant.VendorPerformanceMetric, error) {
	rows, err := r.q.ListVendorPerformanceMetrics(ctx)
	if err != nil {
		return nil, err
	}
	metrics := make([]*merchant.VendorPerformanceMetric, 0, len(rows))
	for _, row := range rows {
		m := toVendorPerformanceMetricDomain(row)
		metrics = append(metrics, &m)
	}
	return metrics, nil
}

func (r *VendorPerformanceMetricRepository) ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*merchant.VendorPerformanceMetric, error) {
	rows, err := r.q.ListVendorPerformanceMetricsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	metrics := make([]*merchant.VendorPerformanceMetric, 0, len(rows))
	for _, row := range rows {
		m := toVendorPerformanceMetricDomain(row)
		metrics = append(metrics, &m)
	}
	return metrics, nil
}

func (r *VendorPerformanceMetricRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVendorPerformanceMetric(ctx, id)
}

func toVendorPerformanceMetricDomain(row sqlc.VendorPerformanceMetric) merchant.VendorPerformanceMetric {
	return merchant.VendorPerformanceMetric{
		ID:         row.ID,
		VendorID:   row.VendorID,
		MetricName: row.MetricName,
		Value:      row.Value,
		RecordedAt: row.RecordedAt,
	}
}
