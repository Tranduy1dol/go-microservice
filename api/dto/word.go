package dto

import "github.com/Tranduy1dol/learning-japanese/internal/domain"

type SearchWordQuery struct {
	Q     string `form:"q" binding:"required"`
	Limit int    `form:"limit" binding:"omitempty,min=1,max=100"`
}

type CreateWordRequest struct {
	Kanji    []KanjiDTO   `json:"kanji" binding:"required,min=1,dive"`
	Readings []ReadingDTO `json:"readings" binding:"required,min=1,dive"`
	Senses   []SenseDTO   `json:"senses" binding:"required,min=1,dive"`
	JLPT     int          `json:"jlpt" binding:"omitempty,min=1,max=5"`
}

type KanjiDTO struct {
	Text string `json:"text" binding:"required"`
	Info string `json:"info"`
}

type ReadingDTO struct {
	Text string `json:"text" binding:"required"`
}

type SenseDTO struct {
	POS   []string   `json:"pos"`
	Gloss []GlossDTO `json:"gloss" binding:"required,min=1,dive"`
}

type GlossDTO struct {
	Text string `json:"text" binding:"required"`
	Lang string `json:"lang" binding:"required"`
}

type WordResponse struct {
	ID       string            `json:"id"`
	Kanji    []KanjiResponse   `json:"kanji"`
	Readings []ReadingResponse `json:"readings"`
	Sense    []SenseResponse   `json:"sense"`
	JLPT     int               `json:"jlpt"`
	IsCommon bool              `json:"is_common"`
}

type KanjiResponse struct {
	Text string `json:"text"`
	Info string `json:"info"`
}

type ReadingResponse struct {
	Text string   `json:"text"`
	Info []string `json:"info"`
}

type SenseResponse struct {
	POS   []string        `json:"pos"`
	Gloss []GlossResponse `json:"gloss"`
}

type GlossResponse struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

func (r *CreateWordRequest) ToDomain() *domain.Word {
	kanji := make([]domain.Kanji, len(r.Kanji))
	for i, k := range r.Kanji {
		kanji[i] = domain.Kanji{
			Text: k.Text,
			Info: k.Info,
		}
	}

	readings := make([]domain.Reading, len(r.Readings))
	for i, rd := range r.Readings {
		readings[i] = domain.Reading{
			Text: rd.Text,
		}
	}

	senses := make([]domain.Sense, len(r.Senses))
	for i, s := range r.Senses {
		glosses := make([]domain.Gloss, len(s.Gloss))
		for j, g := range s.Gloss {
			glosses[j] = domain.Gloss{
				Text: g.Text,
				Lang: g.Lang,
			}
		}

		senses[i] = domain.Sense{
			POS:   s.POS,
			Gloss: glosses,
		}
	}

	return &domain.Word{
		Kanji:    kanji,
		Readings: readings,
		Senses:   senses,
		JLPT:     r.JLPT,
		IsCommon: false,
		Source:   "admin",
	}
}

func NewWordResponse(w *domain.Word) WordResponse {
	kanji := make([]KanjiResponse, len(w.Kanji))
	for i, k := range w.Kanji {
		kanji[i] = KanjiResponse{
			Text: k.Text,
			Info: k.Info,
		}
	}

	readings := make([]ReadingResponse, len(w.Readings))
	for i, r := range w.Readings {
		readings[i] = ReadingResponse{
			Text: r.Text,
			Info: r.Info,
		}
	}

	senses := make([]SenseResponse, len(w.Senses))
	for i, s := range w.Senses {
		gloss := make([]GlossResponse, len(s.Gloss))
		for j, g := range s.Gloss {
			gloss[j] = GlossResponse{
				Text: g.Text,
				Lang: g.Lang,
			}
		}

		senses[i] = SenseResponse{
			POS:   s.POS,
			Gloss: gloss,
		}
	}

	return WordResponse{
		ID:       w.ID,
		JLPT:     w.JLPT,
		IsCommon: w.IsCommon,
		Kanji:    kanji,
		Readings: readings,
		Sense:    senses,
	}
}

func NewWordListResponse(ws []*domain.Word) []WordResponse {
	res := make([]WordResponse, len(ws))
	for i, w := range ws {
		res[i] = NewWordResponse(w)
	}

	return res
}
