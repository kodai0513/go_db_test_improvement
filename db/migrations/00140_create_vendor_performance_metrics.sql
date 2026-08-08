-- +goose Up
CREATE TABLE vendor_performance_metrics (
    id UUID PRIMARY KEY,
    vendor_id UUID NOT NULL,
    metric_name TEXT NOT NULL,
    value INTEGER NOT NULL,
    recorded_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE vendor_performance_metrics;
