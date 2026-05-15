package usecase

import (
	"context"
	"log"

	"github.com/Tranduy1dol/kotoba-press-core/internal/adapter/grpc"
	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"github.com/Tranduy1dol/kotoba-press-core/internal/port"
)

type LookupService struct {
	dictRepo     port.DictionaryRepository
	grammarRepo  port.GrammarRepository
	searchClient *grpc.SearchClient
}

func NewLookupService(dict port.DictionaryRepository, grammar port.GrammarRepository, searchClient *grpc.SearchClient) *LookupService {
	return &LookupService{
		dictRepo:     dict,
		grammarRepo:  grammar,
		searchClient: searchClient,
	}
}

func (s *LookupService) GetWord(ctx context.Context, id string) (*domain.Word, error) {
	return s.dictRepo.GetByID(ctx, id)
}

func (s *LookupService) SearchWord(ctx context.Context, query string, limit int) ([]*domain.Word, error) {
	if s.searchClient != nil {
		resp, err := s.searchClient.Search(ctx, query, limit)
		if err == nil && len(resp.Hits) > 0 {
			words := make([]*domain.Word, 0, len(resp.Hits))
			for _, hit := range resp.Hits {
				word, err := s.dictRepo.GetByID(ctx, hit.DocId)
				if err == nil {
					words = append(words, word)
				}
			}

			if len(words) > 0 {
				return words, nil
			}
		}

		log.Printf("[WARN] groc search failed: %v", err)
	}

	if limit <= 0 || limit > 20 {
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
