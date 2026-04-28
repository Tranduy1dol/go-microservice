package port

import "context"

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByGoogleID(ctx context.Context, googleID string) (*domain.User, error)
	Upsert(ctx context.Context, user *domain.User) error
	UpdateProgress(ctx context.Context, userID string, progress domain.Progress) error
}
