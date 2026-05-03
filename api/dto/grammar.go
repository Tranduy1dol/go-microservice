package dto

import "github.com/Tranduy1dol/learning-japanese/internal/domain"

type CreateGrammarRequest struct {
	Pattern   string              `json:"pattern" binding:"required,min=1"`
	Meaning   string              `json:"meaning" binding:"required,min=1"`
	Formation string              `json:"formation" binding:"required,min=1"`
	JLPT      int                 `json:"jlpt" binding:"required,min=1,max=5"`
	Examples  []GrammarExampleDTO `json:"examples" binding:"omitempty,dive"`
	Notes     string              `json:"notes" binding:"omitempty"`
}

type GrammarExampleDTO struct {
	Japanese    string `json:"japanese" binding:"required"`
	Reading     string `json:"reading" binding:"required"`
	Translation string `json:"translation" binding:"required"`
}

func (r *CreateGrammarRequest) ToDomain() *domain.Grammar {
	examples := make([]domain.GrammarExample, len(r.Examples))
	for i, e := range r.Examples {
		examples[i] = domain.GrammarExample{
			Japanese:    e.Japanese,
			Reading:     e.Reading,
			Translation: e.Translation,
		}
	}

	return &domain.Grammar{
		Pattern:   r.Pattern,
		Meaning:   r.Meaning,
		Formation: r.Formation,
		JLPT:      r.JLPT,
		Example:   examples,
		Notes:     r.Notes,
		Source:    "admin",
	}
}
