package port

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByGoogleID(ctx context.Context, googleID string) (*domain.User, error)
	Upsert(ctx context.Context, googleID, email, name, pictureURL string) (*domain.User, error)
	UpdateProgress(ctx context.Context, userID string, progress domain.Progress) error
}
