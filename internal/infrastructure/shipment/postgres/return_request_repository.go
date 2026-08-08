package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres/sqlc"
)

// ReturnRequestRepository は shipment.ReturnRequestRepository の実装。
type ReturnRequestRepository struct {
	q *sqlc.Queries
}

func NewReturnRequestRepository(db *sql.DB) *ReturnRequestRepository {
	return &ReturnRequestRepository{q: sqlc.New(db)}
}

var _ shipment.ReturnRequestRepository = (*ReturnRequestRepository)(nil)

func (r *ReturnRequestRepository) Create(ctx context.Context, rr *shipment.ReturnRequest) error {
	row, err := r.q.CreateReturnRequest(ctx, sqlc.CreateReturnRequestParams{
		ID:        rr.ID,
		OrderID:   rr.OrderID,
		Reason:    rr.Reason,
		Status:    rr.Status,
		CreatedAt: rr.CreatedAt,
	})
	if err != nil {
		return err
	}
	*rr = toReturnRequestDomain(row)
	return nil
}

func (r *ReturnRequestRepository) Get(ctx context.Context, id uuid.UUID) (*shipment.ReturnRequest, error) {
	row, err := r.q.GetReturnRequest(ctx, id)
	if err != nil {
		return nil, err
	}
	rr := toReturnRequestDomain(row)
	return &rr, nil
}

func (r *ReturnRequestRepository) List(ctx context.Context) ([]*shipment.ReturnRequest, error) {
	rows, err := r.q.ListReturnRequests(ctx)
	if err != nil {
		return nil, err
	}
	requests := make([]*shipment.ReturnRequest, 0, len(rows))
	for _, row := range rows {
		rr := toReturnRequestDomain(row)
		requests = append(requests, &rr)
	}
	return requests, nil
}

func (r *ReturnRequestRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*shipment.ReturnRequest, error) {
	rows, err := r.q.ListReturnRequestsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	requests := make([]*shipment.ReturnRequest, 0, len(rows))
	for _, row := range rows {
		rr := toReturnRequestDomain(row)
		requests = append(requests, &rr)
	}
	return requests, nil
}

func (r *ReturnRequestRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReturnRequest(ctx, id)
}

func toReturnRequestDomain(row sqlc.ReturnRequest) shipment.ReturnRequest {
	return shipment.ReturnRequest{
		ID:        row.ID,
		OrderID:   row.OrderID,
		Reason:    row.Reason,
		Status:    row.Status,
		CreatedAt: row.CreatedAt,
	}
}
