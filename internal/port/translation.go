package port

import "context"

type Translator interface {
	Translate(ctx context.Context, text, sourceLang, targetLang string) (string, error)
}
