package mongo

import (
	"errors"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func wrapError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return domain.ErrNotFound
	}

	if mongo.IsDuplicateKeyError(err) {
		return domain.ErrAlreadyExists
	}

	return err
}
