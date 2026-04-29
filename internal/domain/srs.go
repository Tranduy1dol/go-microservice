package domain

import (
	"math"
	"time"
)

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
	if quality < 3 {
		card.Repetition = 0
		card.Interval = 1
	} else {
		switch card.Repetition {
		case 0:
			card.Interval = 1
		case 1:
			card.Interval = 6
		default:
			card.Interval = int(math.Round(float64(card.Interval) * card.EaseFactor))
		}
		card.Repetition += 1
	}

	card.EaseFactor = card.EaseFactor + (0.1 - float64(5-quality)*(0.08+float64(5-quality)*0.02))
	if card.EaseFactor < 1.3 {
		card.EaseFactor = 1.3
	}

	card.DueDate = time.Now().AddDate(0, 0, card.Interval)
	return card
}
