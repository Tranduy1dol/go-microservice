package domain

import "time"

type SRSCard struct {
	ID         string    `bson:"_id"`
	UserID     string    `bson:"user_id"`
	WordID     string    `bson:"word_id"`
	EaseFactor float64   `bson:"ease_factor"`
	Interval   int       `bson:"interval"`
	Repetition int       `bson:"repetition"`
	DueDate    time.Time `bson:"due_date"`
	CreatedAt  time.Time `bson:"created_at"`
}

func CalculateSM2(quality int, card SRSCard) SRSCard {
}
