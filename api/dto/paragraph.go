package dto

import "github.com/Tranduy1dol/learning-japanese/internal/domain"

type CreateParagraphRequest struct {
	Title     string                 `json:"title" binding:"required,min=1"`
	Content   string                 `json:"content" binding:"required,min=10"`
	JLPT      int                    `json:"jlpt" binding:"required,min=1,max=5"`
	Questions []ParagraphQuestionDto `json:"questions" binding:"required,min=1,dive"`
	Tags      []string               `json:"tags" binding:"omitempty,dive,required"`
}

type ParagraphQuestionDto struct {
	Type         string   `json:"type" binding:"required,oneof=multiple_choice fill_in_blank"`
	JLPT         int      `json:"jlpt" binding:"required,min=1,max=5"`
	Prompt       string   `json:"prompt" binding:"required,min=1"`
	Choices      []string `json:"choices" binding:"required,min=2,max=6,dive,required"`
	CorrectIndex int      `json:"correct_index" binding:"min=0"`
	Explanation  string   `json:"explanation" binding:"omitempty"`
}

func (r *CreateParagraphRequest) ToDomain() *domain.Paragraph {
	questions := make([]domain.Question, len(r.Questions))
	for i, q := range r.Questions {
		questions[i] = q.ToDomain()
	}

	return &domain.Paragraph{
		Title:     r.Title,
		Content:   r.Content,
		JLPT:      r.JLPT,
		Questions: questions,
		Tags:      r.Tags,
		Source:    "admin",
	}
}

func (q *ParagraphQuestionDto) ToDomain() domain.Question {
	return domain.Question{
		Type:         domain.QuestionType(q.Type),
		Section:      domain.SectionReading,
		JLPT:         q.JLPT,
		Prompt:       q.Prompt,
		Choices:      q.Choices,
		CorrectIndex: q.CorrectIndex,
		Explanation:  q.Explanation,
		Source:       "admin",
	}
}
