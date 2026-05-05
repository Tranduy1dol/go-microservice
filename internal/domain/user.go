package domain

import "time"

type User struct {
	ID            string    `bson:"_id"`
	Email         string    `bson:"email"`
	Name          string    `bson:"name"`
	PictureURL    string    `bson:"picture_url"`
	Role          string    `bson:"role"`
	GoogleID      string    `bson:"google_id"`
	CreatedAt     time.Time `bson:"created_at"`
	StudyProgress Progress  `bson:"study_progress"`
}

type ActiviyLog struct {
	Date   string `bson:"date"`
	Action string `bson:"action"`
	Detail string `bson:"detail"`
}

type Progress struct {
	JLPTLevel      int          `bson:"jlpt_level"`
	CardsStudied   int          `bson:"cards_studied" `
	LastStudyAt    time.Time    `bson:"last_study_at" `
	Streak         int          `bson:"streak"`
	LongestStreak  int          `bson:"longest_streak"`
	QuizzesTaken   int          `bson:"quizzes_taken"`
	Accuracy       float64      `bson:"accuracy"`
	WeeklyMinutes  []int        `bson:"weekly_minutes"`
	RecentActivity []ActiviyLog `bson:"recent_activity"`
}
