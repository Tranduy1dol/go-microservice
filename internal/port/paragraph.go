package port

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
)

type ParagraphRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Paragraph, error)
	GetByJLPT(ctx context.Context, level int, limit, offset int) ([]*domain.Paragraph, error)
	GetRandom(ctx context.Context, count int) ([]*domain.Paragraph, error)
	Create(ctx context.Context, p *domain.Paragraph) error
	BulkCreate(ctx context.Context, paragraphs []*domain.Paragraph) (int64, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*domain.Paragraph, int, error)
	Update(ctx context.Context, id string, paragraph *domain.Paragraph) error
}
