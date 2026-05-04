package usecase

import (
	"context"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
	"github.com/Tranduy1dol/learning-japanese/internal/port"
)

type UserService struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetMe(ctx context.Context, userID string) (*domain.User, error) {
	return s.userRepo.GetByID(ctx, userID)
}
