package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres/sqlc"
)

// StockReservationRepository は inventory.StockReservationRepository の実装。
type StockReservationRepository struct {
	q *sqlc.Queries
}

func NewStockReservationRepository(db *sql.DB) *StockReservationRepository {
	return &StockReservationRepository{q: sqlc.New(db)}
}

var _ inventory.StockReservationRepository = (*StockReservationRepository)(nil)

func (r *StockReservationRepository) Create(ctx context.Context, s *inventory.StockReservation) error {
	row, err := r.q.CreateStockReservation(ctx, sqlc.CreateStockReservationParams{
		ID:          s.ID,
		InventoryID: s.InventoryID,
		OrderID:     s.OrderID,
		Quantity:    s.Quantity,
		CreatedAt:   s.CreatedAt,
	})
	if err != nil {
		return err
	}
	*s = toStockReservationDomain(row)
	return nil
}

func (r *StockReservationRepository) Get(ctx context.Context, id uuid.UUID) (*inventory.StockReservation, error) {
	row, err := r.q.GetStockReservation(ctx, id)
	if err != nil {
		return nil, err
	}
	s := toStockReservationDomain(row)
	return &s, nil
}

func (r *StockReservationRepository) List(ctx context.Context) ([]*inventory.StockReservation, error) {
	rows, err := r.q.ListStockReservations(ctx)
	if err != nil {
		return nil, err
	}
	reservations := make([]*inventory.StockReservation, 0, len(rows))
	for _, row := range rows {
		s := toStockReservationDomain(row)
		reservations = append(reservations, &s)
	}
	return reservations, nil
}

func (r *StockReservationRepository) ListByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]*inventory.StockReservation, error) {
	rows, err := r.q.ListStockReservationsByInventoryID(ctx, inventoryID)
	if err != nil {
		return nil, err
	}
	reservations := make([]*inventory.StockReservation, 0, len(rows))
	for _, row := range rows {
		s := toStockReservationDomain(row)
		reservations = append(reservations, &s)
	}
	return reservations, nil
}

func (r *StockReservationRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*inventory.StockReservation, error) {
	rows, err := r.q.ListStockReservationsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	reservations := make([]*inventory.StockReservation, 0, len(rows))
	for _, row := range rows {
		s := toStockReservationDomain(row)
		reservations = append(reservations, &s)
	}
	return reservations, nil
}

func (r *StockReservationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteStockReservation(ctx, id)
}

func toStockReservationDomain(row sqlc.StockReservation) inventory.StockReservation {
	return inventory.StockReservation{
		ID:          row.ID,
		InventoryID: row.InventoryID,
		OrderID:     row.OrderID,
		Quantity:    row.Quantity,
		CreatedAt:   row.CreatedAt,
	}
}
