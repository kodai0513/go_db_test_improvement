package staff

import (
	"context"

	"github.com/google/uuid"
)

// Department は部署を表す。
type Department struct {
	ID   uuid.UUID
	Name string
}

// DepartmentRepository は departments テーブルへのアクセスを抽象化する。
type DepartmentRepository interface {
	Create(ctx context.Context, d *Department) error
	Get(ctx context.Context, id uuid.UUID) (*Department, error)
	List(ctx context.Context) ([]*Department, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
