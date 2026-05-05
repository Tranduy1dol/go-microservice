package dto

// These concrete types exist purely to help Swaggo parse pagination responses
// since some versions of Swaggo crash when parsing Go Generics (PaginatedResponse[T]).

type SwaggerPaginatedWords struct {
	Items  []WordResponse `json:"items"`
	Total  int            `json:"total"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
}

type SwaggerPaginatedGrammars struct {
	Items  []GrammarResponse `json:"items"`
	Total  int               `json:"total"`
	Limit  int               `json:"limit"`
	Offset int               `json:"offset"`
}

type SwaggerPaginatedParagraphs struct {
	Items  []ParagraphResponse `json:"items"`
	Total  int                 `json:"total"`
	Limit  int                 `json:"limit"`
	Offset int                 `json:"offset"`
}

type SwaggerPaginatedQuestions struct {
	Items  []QuestionWithAnswerResponse `json:"items"`
	Total  int                          `json:"total"`
	Limit  int                          `json:"limit"`
	Offset int                          `json:"offset"`
}
