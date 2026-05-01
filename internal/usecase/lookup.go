package usecase

import (
	"context"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
	"github.com/Tranduy1dol/learning-japanese/internal/port"
)

type LookupService struct {
	dictRepo    port.DictionaryRepository
	grammarRepo port.GrammarRepository
}

func NewLookupService(dict port.DictionaryRepository, grammar port.GrammarRepository) *LookupService {
	return &LookupService{
		dictRepo:    dict,
		grammarRepo: grammar,
	}
}

func (s *LookupService) GetWord(ctx context.Context, id string) (*domain.Word, error) {
	return s.dictRepo.GetByID(ctx, id)
}

func (s *LookupService) SearchWord(ctx context.Context, query string, limit int) ([]*domain.Word, error) {
	if limit <= 20 {
		limit = 20
	}
	return s.dictRepo.Search(ctx, query, limit)
}

func (s *LookupService) GetGrammar(ctx context.Context, id string) (*domain.Grammar, error) {
	return s.grammarRepo.GetByID(ctx, id)
}

func (s *LookupService) ListGrammarByJLPT(ctx context.Context, level, limit int) ([]*domain.Grammar, error) {
	return s.grammarRepo.GetByJLPT(ctx, level, limit)
}

func (s *LookupService) BrowseWordByJLPT(ctx context.Context, level, limit, offset int) ([]*domain.Word, int, error) {
	return s.dictRepo.GetByJLPT(ctx, level, limit, offset)
}

func (s *LookupService) LookupAny(ctx context.Context, id string) (any, string, error) {
	if word, err := s.dictRepo.GetByID(ctx, id); err == nil {
		return word, "word", nil
	}

	if grammar, err := s.grammarRepo.GetByID(ctx, id); err == nil {
		return grammar, "grammar", nil
	}

	return nil, "", domain.ErrNotFound
}
