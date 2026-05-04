package dto

import "github.com/Tranduy1dol/kotoba-press-core/internal/domain"

type CreateQuestionRequest struct {
	Type         string   `json:"type" binding:"required,oneof=multiple_choice fill_in_blank reorder"`
	Section      string   `json:"section" binding:"required,oneof=vocabulary grammar reading"`
	JLPT         int      `json:"jlpt" binding:"required,min=1,max=5"`
	Prompt       string   `json:"prompt" binding:"required,min=1"`
	Choices      []string `json:"choices" binding:"required,min=2,max=6,dive,required"`
	CorrectIndex int      `json:"correct_index" binding:"min=0"`
	Explanation  string   `json:"explanation" binding:"omitempty"`
	Tags         []string `json:"tags" binding:"omitempty,dive,required"`
}

type QuestionResponse struct {
	ID      string   `json:"id"`
	Type    string   `json:"type"`
	Prompt  string   `json:"prompt"`
	Choices []string `json:"choices"`
}

type QuestionWithAnswerResponse struct {
	QuestionResponse
	CorrectIndex int    `json:"correct_index"`
	Explanation  string `json:"explanation"`
}

func (r *CreateQuestionRequest) ToDomain() *domain.Question {
	return &domain.Question{
		Type:         domain.QuestionType(r.Type),
		Section:      domain.TestSection(r.Section),
		JLPT:         r.JLPT,
		Prompt:       r.Prompt,
		Choices:      r.Choices,
		CorrectIndex: r.CorrectIndex,
		Explanation:  r.Explanation,
		Tags:         r.Tags,
		Source:       "admin",
	}
}

func NewQuestionResponse(q *domain.Question) QuestionResponse {
	return QuestionResponse{
		ID:      q.ID,
		Type:    string(q.Type),
		Prompt:  q.Prompt,
		Choices: q.Choices,
	}
}

func NewQuestionWithAnswerResponse(q *domain.Question) QuestionWithAnswerResponse {
	questionRes := NewQuestionResponse(q)
	return QuestionWithAnswerResponse{
		QuestionResponse: questionRes,
		CorrectIndex:     q.CorrectIndex,
		Explanation:      q.Explanation,
	}
}
