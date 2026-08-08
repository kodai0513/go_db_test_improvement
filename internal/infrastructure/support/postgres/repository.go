package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres/sqlc"
)

// Repository は support.TicketRepository の実装。
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ support.TicketRepository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, t *support.Ticket) error {
	row, err := r.q.CreateTicket(ctx, sqlc.CreateTicketParams{
		ID:        t.ID,
		MemberID:  t.MemberID,
		Subject:   t.Subject,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
	})
	if err != nil {
		return err
	}
	*t = toTicketDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*support.Ticket, error) {
	row, err := r.q.GetTicket(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toTicketDomain(row)
	return &t, nil
}

func (r *Repository) List(ctx context.Context) ([]*support.Ticket, error) {
	rows, err := r.q.ListTickets(ctx)
	if err != nil {
		return nil, err
	}
	tickets := make([]*support.Ticket, 0, len(rows))
	for _, row := range rows {
		t := toTicketDomain(row)
		tickets = append(tickets, &t)
	}
	return tickets, nil
}

func (r *Repository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*support.Ticket, error) {
	rows, err := r.q.ListTicketsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	tickets := make([]*support.Ticket, 0, len(rows))
	for _, row := range rows {
		t := toTicketDomain(row)
		tickets = append(tickets, &t)
	}
	return tickets, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteTicket(ctx, id)
}

func toTicketDomain(row sqlc.Ticket) support.Ticket {
	return support.Ticket{
		ID:        row.ID,
		MemberID:  row.MemberID,
		Subject:   row.Subject,
		Status:    row.Status,
		CreatedAt: row.CreatedAt,
	}
}
