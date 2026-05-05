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

type ActivityLogResponse struct {
	Date   string `json:"date"`
	Action string `json:"action"`
	Detail string `json:"detail"`
}

type ProgressResponse struct {
	JLPTLevel      int                   `json:"jlpt_level"`
	CardsStudied   int                   `json:"cards_studied"`
	LastStudyAt    time.Time             `json:"last_study_at"`
	Streak         int                   `json:"streak"`
	LongestStreak  int                   `json:"longest_streak"`
	QuizzesTaken   int                   `json:"quizzes_taken"`
	Accuracy       float64               `json:"accuracy"`
	WeeklyMinutes  []int                 `json:"weekly_minutes"`
	RecentActivity []ActivityLogResponse `json:"recent_activity"`
}

func NewUserResponse(u *domain.User) UserResponse {
	recent := make([]ActivityLogResponse, len(u.StudyProgress.RecentActivity))
	for i, log := range u.StudyProgress.RecentActivity {
		recent[i] = ActivityLogResponse{
			Date:   log.Date,
			Action: log.Action,
			Detail: log.Detail,
		}
	}

	return UserResponse{
		ID:         u.ID,
		Email:      u.Email,
		Name:       u.Name,
		PictureURL: u.PictureURL,
		CreatedAt:  u.CreatedAt,
		StudyProgress: ProgressResponse{
			JLPTLevel:      u.StudyProgress.JLPTLevel,
			CardsStudied:   u.StudyProgress.CardsStudied,
			LastStudyAt:    u.StudyProgress.LastStudyAt,
			Streak:         u.StudyProgress.Streak,
			LongestStreak:  u.StudyProgress.LongestStreak,
			QuizzesTaken:   u.StudyProgress.QuizzesTaken,
			Accuracy:       u.StudyProgress.Accuracy,
			WeeklyMinutes:  u.StudyProgress.WeeklyMinutes,
			RecentActivity: recent,
		},
	}
}
