package port

import (
	"context"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
)

type SRSRepository interface {
	Save(ctx context.Context, card *domain.SRSCard) error
	GetDueCards(ctx context.Context, userID string, limit int) ([]*domain.SRSCard, error)
	GetByIDAndUser(ctx context.Context, id, userID string) (*domain.SRSCard, error)
	GetByWordAndUser(ctx context.Context, wordID, userID string) (*domain.SRSCard, error)
}
