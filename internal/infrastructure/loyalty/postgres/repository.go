package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres/sqlc"
)

// Repository は loyalty.LoyaltyProgramRepository の実装。
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ loyalty.LoyaltyProgramRepository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, p *loyalty.LoyaltyProgram) error {
	row, err := r.q.CreateLoyaltyProgram(ctx, sqlc.CreateLoyaltyProgramParams{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
	})
	if err != nil {
		return err
	}
	*p = toLoyaltyProgramDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*loyalty.LoyaltyProgram, error) {
	row, err := r.q.GetLoyaltyProgram(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toLoyaltyProgramDomain(row)
	return &p, nil
}

func (r *Repository) List(ctx context.Context) ([]*loyalty.LoyaltyProgram, error) {
	rows, err := r.q.ListLoyaltyPrograms(ctx)
	if err != nil {
		return nil, err
	}
	programs := make([]*loyalty.LoyaltyProgram, 0, len(rows))
	for _, row := range rows {
		p := toLoyaltyProgramDomain(row)
		programs = append(programs, &p)
	}
	return programs, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteLoyaltyProgram(ctx, id)
}

func toLoyaltyProgramDomain(row sqlc.LoyaltyProgram) loyalty.LoyaltyProgram {
	return loyalty.LoyaltyProgram{
		ID:        row.ID,
		Name:      row.Name,
		CreatedAt: row.CreatedAt,
	}
}
