package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"github.com/Tranduy1dol/kotoba-press-core/internal/port"
	"github.com/google/uuid"
)

type SRSService struct {
	srsRepo  port.SRSRepository
	wordRepo port.DictionaryRepository
}

func NewSRSService(srsRepo port.SRSRepository, wordRepo port.DictionaryRepository) *SRSService {
	return &SRSService{
		srsRepo:  srsRepo,
		wordRepo: wordRepo,
	}
}

func (s *SRSService) AddWordToDeck(ctx context.Context, userID, wordID string) (*domain.SRSCard, error) {
	_, err := s.wordRepo.GetByID(ctx, wordID)
	if err != nil {
		return nil, err
	}

	existingCard, err := s.srsRepo.GetByWordAndUser(ctx, wordID, userID)
	if err == nil && existingCard != nil {
		return nil, domain.ErrAlreadyExists
	}

	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return nil, err
	}

	newCard := &domain.SRSCard{
		ID:         uuid.New().String(),
		UserID:     userID,
		WordID:     wordID,
		EaseFactor: 2.5,
		Interval:   0,
		Repetition: 0,
		DueDate:    time.Now(),
		CreatedAt:  time.Now(),
	}

	if err := s.srsRepo.Save(ctx, newCard); err != nil {
		return nil, err
	}

	return newCard, nil
}

func (s *SRSService) GetDueCards(ctx context.Context, userID string, limit int) ([]*domain.SRSCard, error) {
	if limit <= 0 {
		limit = 20
	}

	return s.srsRepo.GetDueCards(ctx, userID, limit)
}

func (s *SRSService) ReviewCard(ctx context.Context, userID, cardID string, quality int) (*domain.SRSCard, error) {
	if quality < 0 || quality > 5 {
		return nil, domain.ErrInvalidInput
	}

	card, err := s.srsRepo.GetByIDAndUser(ctx, cardID, userID)
	if err != nil {
		return nil, err
	}

	*card = domain.CalculateSM2(quality, *card)
	if err := s.srsRepo.Save(ctx, card); err != nil {
		return nil, err
	}

	return card, nil
}
