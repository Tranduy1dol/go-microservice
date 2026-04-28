package domain

type QuestionType string

const (
	MultipleChoie QuestionType = "multiple_choice"
	FillInBlank   QuestionType = "fill_in_blank"
	Reorder       QuestionType = "reorder"
)

type Question struct {
	ID           string       `bson:"_id"`
	Type         QuestionType `bson:"type"`
	Section      TestSection  `bson:"section"`
	JLPT         int          `bson:"jlpt"`
	Promt        string       `bson:"promt"`
	Choices      []string     `bson:"choices"`
	CorrectIndex int          `bson:"correct_index"`
	Explanation  string       `bson:"explanation"`
	Tags         []string     `bson:"tags"`
	Source       string       `bson:"source"`
}
