package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres/sqlc"
)

// SatisfactionSurveyRepository は support.SatisfactionSurveyRepository の実装。
type SatisfactionSurveyRepository struct {
	q *sqlc.Queries
}

func NewSatisfactionSurveyRepository(db *sql.DB) *SatisfactionSurveyRepository {
	return &SatisfactionSurveyRepository{q: sqlc.New(db)}
}

var _ support.SatisfactionSurveyRepository = (*SatisfactionSurveyRepository)(nil)

func (r *SatisfactionSurveyRepository) Create(ctx context.Context, s *support.SatisfactionSurvey) error {
	row, err := r.q.CreateSatisfactionSurvey(ctx, sqlc.CreateSatisfactionSurveyParams{
		ID:          s.ID,
		TicketID:    s.TicketID,
		Score:       s.Score,
		Comment:     s.Comment,
		SubmittedAt: s.SubmittedAt,
	})
	if err != nil {
		return err
	}
	*s = toSatisfactionSurveyDomain(row)
	return nil
}

func (r *SatisfactionSurveyRepository) Get(ctx context.Context, id uuid.UUID) (*support.SatisfactionSurvey, error) {
	row, err := r.q.GetSatisfactionSurvey(ctx, id)
	if err != nil {
		return nil, err
	}
	s := toSatisfactionSurveyDomain(row)
	return &s, nil
}

func (r *SatisfactionSurveyRepository) List(ctx context.Context) ([]*support.SatisfactionSurvey, error) {
	rows, err := r.q.ListSatisfactionSurveys(ctx)
	if err != nil {
		return nil, err
	}
	surveys := make([]*support.SatisfactionSurvey, 0, len(rows))
	for _, row := range rows {
		s := toSatisfactionSurveyDomain(row)
		surveys = append(surveys, &s)
	}
	return surveys, nil
}

func (r *SatisfactionSurveyRepository) ListByTicketID(ctx context.Context, ticketID uuid.UUID) ([]*support.SatisfactionSurvey, error) {
	rows, err := r.q.ListSatisfactionSurveysByTicketID(ctx, ticketID)
	if err != nil {
		return nil, err
	}
	surveys := make([]*support.SatisfactionSurvey, 0, len(rows))
	for _, row := range rows {
		s := toSatisfactionSurveyDomain(row)
		surveys = append(surveys, &s)
	}
	return surveys, nil
}

func (r *SatisfactionSurveyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteSatisfactionSurvey(ctx, id)
}

func toSatisfactionSurveyDomain(row sqlc.SatisfactionSurvey) support.SatisfactionSurvey {
	return support.SatisfactionSurvey{
		ID:          row.ID,
		TicketID:    row.TicketID,
		Score:       row.Score,
		Comment:     row.Comment,
		SubmittedAt: row.SubmittedAt,
	}
}
