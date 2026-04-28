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

type Progress struct {
	JLPTLevel    int       `bson:"jlpt_level"`
	CardsStudied int       `bson:"cards_studied"`
	LastStudyAt  time.Time `bson:"last_study_at"`
}
