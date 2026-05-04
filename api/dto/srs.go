package dto

import (
	"time"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
)

type AddWordToDeckRequest struct {
	WordID string `json:"word_id" binding:"required"`
}

type ReviewCardRequest struct {
	Quality int `json:"quality" binding:"required,min=0,max=5"`
}

type SRSCardResponse struct {
	ID         string    `json:"id"`
	WordID     string    `json:"word_id"`
	EaseFactor float64   `json:"ease_factor"`
	Interval   int       `json:"interval"`
	Repetition int       `json:"repetition"`
	DueDate    time.Time `json:"due_date"`
}

func NewSRSCardResponse(card *domain.SRSCard) SRSCardResponse {
	return SRSCardResponse{
		ID:         card.ID,
		WordID:     card.WordID,
		EaseFactor: card.EaseFactor,
		Interval:   card.Interval,
		Repetition: card.Repetition,
		DueDate:    card.DueDate,
	}
}

func NewSRSCardListResponse(cards []*domain.SRSCard) []SRSCardResponse {
	res := make([]SRSCardResponse, len(cards))
	for i, c := range cards {
		res[i] = NewSRSCardResponse(c)
	}

	return res
}
