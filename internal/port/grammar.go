package port

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
)

type GrammarRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Grammar, error)
	GetByJLPT(ctx context.Context, level int, limit int) ([]*domain.Grammar, error)
	GetRandom(ctx context.Context, count int) ([]*domain.Grammar, error)
	Create(ctx context.Context, g *domain.Grammar) error
	BulkCreate(ctx context.Context, grammars []*domain.Grammar) (int64, error)
	Delete(ctx context.Context, id string) error
}
