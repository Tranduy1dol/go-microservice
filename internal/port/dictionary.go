package port

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
)

type DictionaryRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Word, error)
	GetByKanji(ctx context.Context, kanji string) (*domain.Word, error)
	GetByReading(ctx context.Context, reading string) ([]*domain.Word, error)
	Search(ctx context.Context, query string, limit int) ([]*domain.Word, error)
	GetByJLPT(ctx context.Context, level int, limit, offset int) ([]*domain.Word, int, error)
	GetRandom(ctx context.Context, count int) ([]*domain.Word, error)
	Create(ctx context.Context, word *domain.Word) error
	BulkCreate(ctx context.Context, words []*domain.Word) (int64, error)
	Delete(ctx context.Context, id string) error
}
