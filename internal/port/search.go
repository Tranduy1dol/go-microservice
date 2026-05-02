package port

import "context"

type SearchEngine interface {
	Search(ctx context.Context, query string, topK int) ([]SearchResult, error)
}

type SearchResult struct {
	WordID string
	Score  float64
	Title  string
}
