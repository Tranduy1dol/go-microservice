package dto

import "github.com/Tranduy1dol/learning-japanese/internal/domain"

type SubmitTestRequest struct {
	Answers map[string]int `json:"answers" binding:"required"`
}

type SubmitTestResponse struct {
	Score int `json:"score"`
	Total int `json:"total"`
}

type TestResponse struct {
	ID        string             `json:"id"`
	JLPT      int                `json:"jlpt"`
	Sections  []TestPartResponse `json:"sections"`
	TimeLimit int                `json:"time_limit"`
}

type TestPartResponse struct {
	Section   string             `json:"section"`
	Questions []QuestionResponse `json:"questions"`
}

func NewTestResponse(t *domain.Test) TestResponse {
	testPart := make([]TestPartResponse, len(t.Sections))
	for i, p := range t.Sections {
		questionRes := make([]QuestionResponse, len(p.Questions))
		for j, q := range p.Questions {
			questionRes[j] = NewQuestionResponse(&q)
		}

		testPart[i] = TestPartResponse{
			Section:   string(p.Section),
			Questions: questionRes,
		}
	}

	return TestResponse{
		ID:        t.ID,
		JLPT:      t.JLPT,
		Sections:  testPart,
		TimeLimit: t.TimeLimit,
	}
}
