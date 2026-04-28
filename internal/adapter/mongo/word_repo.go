package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type WordRepository struct {
	collection *mongo.Collection
}

func NewWordRepository(db *mongo.Database) *WordRepository {
	return &WordRepository{collection: db.collection("words")}
}

func (r *WordRepository) GetByID(ctx context.Context, id string) (*domain.Word, error) {
}

func (r *WordRepository) GetRandom(ctx context.Context, count int) ([]*domain.Word, error) {
}
