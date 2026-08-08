package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres/sqlc"
)

// ReviewInvitationRepository は review.ReviewInvitationRepository の実装。
type ReviewInvitationRepository struct {
	q *sqlc.Queries
}

func NewReviewInvitationRepository(db *sql.DB) *ReviewInvitationRepository {
	return &ReviewInvitationRepository{q: sqlc.New(db)}
}

var _ review.ReviewInvitationRepository = (*ReviewInvitationRepository)(nil)

func (r *ReviewInvitationRepository) Create(ctx context.Context, i *review.ReviewInvitation) error {
	row, err := r.q.CreateReviewInvitation(ctx, sqlc.CreateReviewInvitationParams{
		ID:       i.ID,
		OrderID:  i.OrderID,
		MemberID: i.MemberID,
		SentAt:   i.SentAt,
	})
	if err != nil {
		return err
	}
	*i = toReviewInvitationDomain(row)
	return nil
}

func (r *ReviewInvitationRepository) Get(ctx context.Context, id uuid.UUID) (*review.ReviewInvitation, error) {
	row, err := r.q.GetReviewInvitation(ctx, id)
	if err != nil {
		return nil, err
	}
	i := toReviewInvitationDomain(row)
	return &i, nil
}

func (r *ReviewInvitationRepository) List(ctx context.Context) ([]*review.ReviewInvitation, error) {
	rows, err := r.q.ListReviewInvitations(ctx)
	if err != nil {
		return nil, err
	}
	invitations := make([]*review.ReviewInvitation, 0, len(rows))
	for _, row := range rows {
		i := toReviewInvitationDomain(row)
		invitations = append(invitations, &i)
	}
	return invitations, nil
}

func (r *ReviewInvitationRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*review.ReviewInvitation, error) {
	rows, err := r.q.ListReviewInvitationsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	invitations := make([]*review.ReviewInvitation, 0, len(rows))
	for _, row := range rows {
		i := toReviewInvitationDomain(row)
		invitations = append(invitations, &i)
	}
	return invitations, nil
}

func (r *ReviewInvitationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReviewInvitation(ctx, id)
}

func toReviewInvitationDomain(row sqlc.ReviewInvitation) review.ReviewInvitation {
	return review.ReviewInvitation{
		ID:       row.ID,
		OrderID:  row.OrderID,
		MemberID: row.MemberID,
		SentAt:   row.SentAt,
	}
}
