package port

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
)

type TestRepository interface {
	Save(ctx context.Context, test *domain.Test) error
	GetByID(ctx context.Context, id string) (*domain.Test, error)
}
