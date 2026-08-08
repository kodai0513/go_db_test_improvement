package support

import (
	"context"

	"github.com/google/uuid"
)

// SupportAgent はサポート担当者を表す。
type SupportAgent struct {
	ID    uuid.UUID
	Name  string
	Email string
}

// SupportAgentRepository は support_agents テーブルへのアクセスを抽象化する。
type SupportAgentRepository interface {
	Create(ctx context.Context, a *SupportAgent) error
	Get(ctx context.Context, id uuid.UUID) (*SupportAgent, error)
	List(ctx context.Context) ([]*SupportAgent, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
