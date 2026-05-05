package port

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
)

type QuestionRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Question, error)
	GetByJLPT(ctx context.Context, level int, section domain.TestSection, limit int) ([]*domain.Question, error)
	GetRandom(ctx context.Context, count int) ([]*domain.Question, error)
	Create(ctx context.Context, q *domain.Question) error
	BulkCreate(ctx context.Context, questions []*domain.Question) (int64, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*domain.Question, int, error)
	Update(ctx context.Context, id string, question *domain.Question) error
}
