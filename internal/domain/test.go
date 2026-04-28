package domain

import "time"

type TestSection string

const (
	SectionVocab   TestSection = "vocabulary"
	SectionGrammar TestSection = "grammar"
	SectionReading TestSection = "reading"
)

type Test struct {
	ID        string     `bson:"_id"`
	JLPT      int        `bson:"jlpt"`
	Sections  []TestPart `bson:"sections"`
	TimeLimit int        `bson:"time_limit"`
	CreatedAt time.Time  `bson:"created_at"`
}

type TestPart struct {
	Section   TestSection `bson:"section"`
	Questions []Question  `bson:"questions"`
}
