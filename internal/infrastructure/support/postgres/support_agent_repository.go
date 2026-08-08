package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres/sqlc"
)

// SupportAgentRepository は support.SupportAgentRepository の実装。
type SupportAgentRepository struct {
	q *sqlc.Queries
}

func NewSupportAgentRepository(db *sql.DB) *SupportAgentRepository {
	return &SupportAgentRepository{q: sqlc.New(db)}
}

var _ support.SupportAgentRepository = (*SupportAgentRepository)(nil)

func (r *SupportAgentRepository) Create(ctx context.Context, a *support.SupportAgent) error {
	row, err := r.q.CreateSupportAgent(ctx, sqlc.CreateSupportAgentParams{
		ID:    a.ID,
		Name:  a.Name,
		Email: a.Email,
	})
	if err != nil {
		return err
	}
	*a = toSupportAgentDomain(row)
	return nil
}

func (r *SupportAgentRepository) Get(ctx context.Context, id uuid.UUID) (*support.SupportAgent, error) {
	row, err := r.q.GetSupportAgent(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toSupportAgentDomain(row)
	return &a, nil
}

func (r *SupportAgentRepository) List(ctx context.Context) ([]*support.SupportAgent, error) {
	rows, err := r.q.ListSupportAgents(ctx)
	if err != nil {
		return nil, err
	}
	agents := make([]*support.SupportAgent, 0, len(rows))
	for _, row := range rows {
		a := toSupportAgentDomain(row)
		agents = append(agents, &a)
	}
	return agents, nil
}

func (r *SupportAgentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteSupportAgent(ctx, id)
}

func toSupportAgentDomain(row sqlc.SupportAgent) support.SupportAgent {
	return support.SupportAgent{
		ID:    row.ID,
		Name:  row.Name,
		Email: row.Email,
	}
}
