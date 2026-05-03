package dto

import "github.com/Tranduy1dol/learning-japanese/internal/domain"

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
