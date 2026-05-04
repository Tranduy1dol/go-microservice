package dto

import (
	"time"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
)

type UserResponse struct {
	ID            string           `json:"id"`
	Email         string           `json:"email"`
	Name          string           `json:"name"`
	PictureURL    string           `json:"picture_url"`
	CreatedAt     time.Time        `json:"created_at"`
	StudyProgress ProgressResponse `json:"study_progress"`
}

type ProgressResponse struct {
	JLPTLevel    int       `json:"jlpt_level"`
	CardsStudied int       `json:"cards_studied"`
	LastStudyAt  time.Time `json:"last_study_at"`
}

func NewUserResponse(u *domain.User) UserResponse {
	return UserResponse{
		ID:         u.ID,
		Email:      u.Email,
		Name:       u.Name,
		PictureURL: u.PictureURL,
		CreatedAt:  u.CreatedAt,
		StudyProgress: ProgressResponse{
			JLPTLevel:    u.StudyProgress.JLPTLevel,
			CardsStudied: u.StudyProgress.CardsStudied,
			LastStudyAt:  u.StudyProgress.LastStudyAt,
		},
	}
}
