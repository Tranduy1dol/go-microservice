package mongo

import (
	"context"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WordRepository struct {
	collection *mongo.Collection
}

func NewWordRepository(db *mongo.Database) *WordRepository {
	return &WordRepository{collection: db.Collection("words")}
}

func (r *WordRepository) GetByID(ctx context.Context, id string) (*domain.Word, error) {
	filter := bson.M{"_id": id}
	var word domain.Word
	err := r.collection.FindOne(ctx, filter).Decode(&word)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return &word, nil
}

func (r *WordRepository) GetRandom(ctx context.Context, count int) ([]*domain.Word, error) {
	filter := bson.M{}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$sample", Value: bson.M{"size": count}}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	var words []*domain.Word
	if err := cursor.All(ctx, &words); err != nil {
		return nil, err
	}
	return words, nil
}
