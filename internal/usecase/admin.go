package usecase

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"github.com/Tranduy1dol/kotoba-press-core/internal/port"
)

type AdminService struct {
	wordRepo      port.DictionaryRepository
	questionRepo  port.QuestionRepository
	paragraphRepo port.ParagraphRepository
	grammarRepo   port.GrammarRepository
}

func NewAdminService(wordRepo port.DictionaryRepository, questionRepo port.QuestionRepository, paragraphRepo port.ParagraphRepository, grammarRepo port.GrammarRepository) *AdminService {
	return &AdminService{
		wordRepo:      wordRepo,
		questionRepo:  questionRepo,
		paragraphRepo: paragraphRepo,
		grammarRepo:   grammarRepo,
	}
}

func (s *AdminService) CreateWord(ctx context.Context, word *domain.Word) error {
	return s.wordRepo.Create(ctx, word)
}

func (s *AdminService) DeleteWord(ctx context.Context, id string) error {
	return s.wordRepo.Delete(ctx, id)
}

func (s *AdminService) ListWords(ctx context.Context, limit, offset int) ([]*domain.Word, int, error) {
	return s.wordRepo.List(ctx, limit, offset)
}

func (s *AdminService) UpdateWord(ctx context.Context, id string, word *domain.Word) error {
	word.ID = id
	return s.wordRepo.Update(ctx, id, word)
}

func (s *AdminService) ListGrammars(ctx context.Context, limit, offset int) ([]*domain.Grammar, int, error) {
	return s.grammarRepo.List(ctx, limit, offset)
}

func (s *AdminService) UpdateGrammar(ctx context.Context, id string, grammar *domain.Grammar) error {
	grammar.ID = id
	return s.grammarRepo.Update(ctx, id, grammar)
}

func (s *AdminService) ListParagraphs(ctx context.Context, limit, offset int) ([]*domain.Paragraph, int, error) {
	return s.paragraphRepo.List(ctx, limit, offset)
}

func (s *AdminService) UpdateParagraph(ctx context.Context, id string, paragraph *domain.Paragraph) error {
	paragraph.ID = id
	return s.paragraphRepo.Update(ctx, id, paragraph)
}

func (s *AdminService) ListQuestions(ctx context.Context, limit, offset int) ([]*domain.Question, int, error) {
	return s.questionRepo.List(ctx, limit, offset)
}

func (s *AdminService) UpdateQuestion(ctx context.Context, id string, question *domain.Question) error {
	question.ID = id
	return s.questionRepo.Update(ctx, id, question)
}

func (s *AdminService) CreateQuestion(ctx context.Context, question *domain.Question) error {
	if question.CorrectIndex >= len(question.Choices) {
		return domain.ErrInvalidInput
	}

	return s.questionRepo.Create(ctx, question)
}

func (s *AdminService) DeleteQuestion(ctx context.Context, id string) error {
	return s.questionRepo.Delete(ctx, id)
}

func (s *AdminService) CreateParagraph(ctx context.Context, paragraph *domain.Paragraph) error {
	return s.paragraphRepo.Create(ctx, paragraph)
}

func (s *AdminService) DeleteParagraph(ctx context.Context, id string) error {
	return s.paragraphRepo.Delete(ctx, id)
}

func (s *AdminService) CreateGrammar(ctx context.Context, grammar *domain.Grammar) error {
	return s.grammarRepo.Create(ctx, grammar)
}

func (s *AdminService) DeleteGrammar(ctx context.Context, id string) error {
	return s.grammarRepo.Delete(ctx, id)
}
