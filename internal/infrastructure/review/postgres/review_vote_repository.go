package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres/sqlc"
)

// ReviewVoteRepository は review.ReviewVoteRepository の実装。
type ReviewVoteRepository struct {
	q *sqlc.Queries
}

func NewReviewVoteRepository(db *sql.DB) *ReviewVoteRepository {
	return &ReviewVoteRepository{q: sqlc.New(db)}
}

var _ review.ReviewVoteRepository = (*ReviewVoteRepository)(nil)

func (r *ReviewVoteRepository) Create(ctx context.Context, v *review.ReviewVote) error {
	row, err := r.q.CreateReviewVote(ctx, sqlc.CreateReviewVoteParams{
		ID:        v.ID,
		ReviewID:  v.ReviewID,
		MemberID:  v.MemberID,
		IsHelpful: v.IsHelpful,
	})
	if err != nil {
		return err
	}
	*v = toReviewVoteDomain(row)
	return nil
}

func (r *ReviewVoteRepository) Get(ctx context.Context, id uuid.UUID) (*review.ReviewVote, error) {
	row, err := r.q.GetReviewVote(ctx, id)
	if err != nil {
		return nil, err
	}
	v := toReviewVoteDomain(row)
	return &v, nil
}

func (r *ReviewVoteRepository) List(ctx context.Context) ([]*review.ReviewVote, error) {
	rows, err := r.q.ListReviewVotes(ctx)
	if err != nil {
		return nil, err
	}
	votes := make([]*review.ReviewVote, 0, len(rows))
	for _, row := range rows {
		v := toReviewVoteDomain(row)
		votes = append(votes, &v)
	}
	return votes, nil
}

func (r *ReviewVoteRepository) ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*review.ReviewVote, error) {
	rows, err := r.q.ListReviewVotesByReviewID(ctx, reviewID)
	if err != nil {
		return nil, err
	}
	votes := make([]*review.ReviewVote, 0, len(rows))
	for _, row := range rows {
		v := toReviewVoteDomain(row)
		votes = append(votes, &v)
	}
	return votes, nil
}

func (r *ReviewVoteRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReviewVote(ctx, id)
}

func toReviewVoteDomain(row sqlc.ReviewVote) review.ReviewVote {
	return review.ReviewVote{
		ID:        row.ID,
		ReviewID:  row.ReviewID,
		MemberID:  row.MemberID,
		IsHelpful: row.IsHelpful,
	}
}
