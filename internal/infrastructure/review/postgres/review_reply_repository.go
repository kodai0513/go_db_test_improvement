package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres/sqlc"
)

// ReviewReplyRepository は review.ReviewReplyRepository の実装。
type ReviewReplyRepository struct {
	q *sqlc.Queries
}

func NewReviewReplyRepository(db *sql.DB) *ReviewReplyRepository {
	return &ReviewReplyRepository{q: sqlc.New(db)}
}

var _ review.ReviewReplyRepository = (*ReviewReplyRepository)(nil)

func (r *ReviewReplyRepository) Create(ctx context.Context, rep *review.ReviewReply) error {
	row, err := r.q.CreateReviewReply(ctx, sqlc.CreateReviewReplyParams{
		ID:        rep.ID,
		ReviewID:  rep.ReviewID,
		MemberID:  rep.MemberID,
		Comment:   rep.Comment,
		CreatedAt: rep.CreatedAt,
	})
	if err != nil {
		return err
	}
	*rep = toReviewReplyDomain(row)
	return nil
}

func (r *ReviewReplyRepository) Get(ctx context.Context, id uuid.UUID) (*review.ReviewReply, error) {
	row, err := r.q.GetReviewReply(ctx, id)
	if err != nil {
		return nil, err
	}
	rep := toReviewReplyDomain(row)
	return &rep, nil
}

func (r *ReviewReplyRepository) List(ctx context.Context) ([]*review.ReviewReply, error) {
	rows, err := r.q.ListReviewReplies(ctx)
	if err != nil {
		return nil, err
	}
	replies := make([]*review.ReviewReply, 0, len(rows))
	for _, row := range rows {
		rep := toReviewReplyDomain(row)
		replies = append(replies, &rep)
	}
	return replies, nil
}

func (r *ReviewReplyRepository) ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*review.ReviewReply, error) {
	rows, err := r.q.ListReviewRepliesByReviewID(ctx, reviewID)
	if err != nil {
		return nil, err
	}
	replies := make([]*review.ReviewReply, 0, len(rows))
	for _, row := range rows {
		rep := toReviewReplyDomain(row)
		replies = append(replies, &rep)
	}
	return replies, nil
}

func (r *ReviewReplyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReviewReply(ctx, id)
}

func toReviewReplyDomain(row sqlc.ReviewReply) review.ReviewReply {
	return review.ReviewReply{
		ID:        row.ID,
		ReviewID:  row.ReviewID,
		MemberID:  row.MemberID,
		Comment:   row.Comment,
		CreatedAt: row.CreatedAt,
	}
}
