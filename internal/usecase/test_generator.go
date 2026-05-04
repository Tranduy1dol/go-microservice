package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"github.com/Tranduy1dol/kotoba-press-core/internal/port"
)

type TestGeneratorService struct {
	questionRepo  port.QuestionRepository
	paragraphRepo port.ParagraphRepository
	testRepo      port.TestRepository
}

func NewTestGeneratorService(question port.QuestionRepository, paragraph port.ParagraphRepository, test port.TestRepository) *TestGeneratorService {
	return &TestGeneratorService{
		questionRepo:  question,
		paragraphRepo: paragraph,
		testRepo:      test,
	}
}

func (s *TestGeneratorService) GenerateTest(ctx context.Context, jlptLevel int) (*domain.Test, error) {
	vocabQs, err := s.questionRepo.GetByJLPT(ctx, jlptLevel, domain.SectionVocab, 25)
	if err != nil {
		return nil, err
	}
	s.shuffleQuestions(vocabQs)

	grammarQs, err := s.questionRepo.GetByJLPT(ctx, jlptLevel, domain.SectionGrammar, 25)
	if err != nil {
		return nil, err
	}
	s.shuffleQuestions(grammarQs)

	paragraphQs, err := s.paragraphRepo.GetRandom(ctx, 3)
	if err != nil {
		return nil, err
	}

	test := &domain.Test{
		ID:   generateID(),
		JLPT: jlptLevel,
		Sections: []domain.TestPart{
			{
				Section:   domain.SectionVocab,
				Questions: s.toQuestionSlice(vocabQs),
			},
			{
				Section:   domain.SectionGrammar,
				Questions: s.toQuestionSlice(grammarQs),
			},
			{
				Section:   domain.SectionReading,
				Questions: s.paragraphsToQuestions(paragraphQs),
			},
		},
		TimeLimit: 95,
		CreatedAt: time.Now(),
	}

	if err := s.testRepo.Save(ctx, test); err != nil {
		return nil, err
	}

	return test, nil
}

func (s *TestGeneratorService) SubmitTest(ctx context.Context, testID string, userAnswer map[string]int) (int, int, error) {
	test, err := s.testRepo.GetByID(ctx, testID)
	if err != nil {
		return 0, 0, err
	}

	totalQuestion := 0
	correctAnswer := 0

	for _, section := range test.Sections {
		for _, q := range section.Questions {
			totalQuestion++
			if userSelected, ok := userAnswer[q.ID]; ok {
				if userSelected == q.CorrectIndex {
					correctAnswer++
				}
			}
		}
	}

	return correctAnswer, totalQuestion, nil
}

func (s *TestGeneratorService) shuffleQuestions(qs []*domain.Question) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(qs), func(i, j int) {
		qs[i], qs[j] = qs[j], qs[i]
	})

	for _, q := range qs {
		correctAnswer := q.Choices[q.CorrectIndex]

		r.Shuffle(len(q.Choices), func(i, j int) {
			q.Choices[i], q.Choices[j] = q.Choices[j], q.Choices[i]
		})

		for i, c := range q.Choices {
			if c == correctAnswer {
				q.CorrectIndex = i
				break
			}
		}
	}
}

func generateID() string {
	return fmt.Sprintf("test-%d", time.Now().UnixNano())
}

func (s *TestGeneratorService) toQuestionSlice(qs []*domain.Question) []domain.Question {
	result := make([]domain.Question, len(qs))
	for i, q := range qs {
		result[i] = *q
	}

	return result
}

func (s *TestGeneratorService) paragraphsToQuestions(paragraphs []*domain.Paragraph) []domain.Question {
	var questions []domain.Question
	for _, p := range paragraphs {
		questions = append(questions, p.Questions...)
	}
	s.shuffleQuestionsFromList(questions)

	return questions
}

func (s *TestGeneratorService) shuffleQuestionsFromList(qs []domain.Question) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(qs), func(i, j int) {
		qs[i], qs[j] = qs[j], qs[i]
	})
}
