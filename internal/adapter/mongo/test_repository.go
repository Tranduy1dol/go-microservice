package mongo

import (
	"context"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestRepository struct {
	collection *mongo.Collection
}

func NewTestRepository(db *mongo.Database) *TestRepository {
	return &TestRepository{
		collection: db.Collection("tests"),
	}
}

func (r *TestRepository) Save(ctx context.Context, test *domain.Test) error {
	_, err := r.collection.InsertOne(ctx, test)
	return wrapError(err)
}

func (r *TestRepository) GetByID(ctx context.Context, id string) (*domain.Test, error) {
	filter := bson.M{"_id": id}

	var test domain.Test
	err := r.collection.FindOne(ctx, filter).Decode(&test)
	return &test, wrapError(err)
}
