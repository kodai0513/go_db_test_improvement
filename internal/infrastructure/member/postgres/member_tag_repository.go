package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// MemberTagRepository は member.MemberTagRepository の実装。
type MemberTagRepository struct {
	q *sqlc.Queries
}

func NewMemberTagRepository(db *sql.DB) *MemberTagRepository {
	return &MemberTagRepository{q: sqlc.New(db)}
}

var _ member.MemberTagRepository = (*MemberTagRepository)(nil)

func (r *MemberTagRepository) Create(ctx context.Context, t *member.MemberTag) error {
	row, err := r.q.CreateMemberTag(ctx, sqlc.CreateMemberTagParams{
		ID:        t.ID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
	})
	if err != nil {
		return err
	}
	*t = toMemberTagDomain(row)
	return nil
}

func (r *MemberTagRepository) Get(ctx context.Context, id uuid.UUID) (*member.MemberTag, error) {
	row, err := r.q.GetMemberTag(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toMemberTagDomain(row)
	return &t, nil
}

func (r *MemberTagRepository) List(ctx context.Context) ([]*member.MemberTag, error) {
	rows, err := r.q.ListMemberTags(ctx)
	if err != nil {
		return nil, err
	}
	tags := make([]*member.MemberTag, 0, len(rows))
	for _, row := range rows {
		t := toMemberTagDomain(row)
		tags = append(tags, &t)
	}
	return tags, nil
}

func (r *MemberTagRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMemberTag(ctx, id)
}

func toMemberTagDomain(row sqlc.MemberTag) member.MemberTag {
	return member.MemberTag{
		ID:        row.ID,
		Name:      row.Name,
		CreatedAt: row.CreatedAt,
	}
}
