package mongo

import (
	"context"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GrammarRepository struct {
	collection *mongo.Collection
}

func NewGrammarRepository(db *mongo.Database) *GrammarRepository {
	return &GrammarRepository{collection: db.Collection("grammar")}
}

func (r *GrammarRepository) GetByID(ctx context.Context, id string) (*domain.Grammar, error) {
	filter := bson.M{"_id": id}

	var grammar domain.Grammar
	err := r.collection.FindOne(ctx, filter).Decode(&grammar)
	return &grammar, wrapError(err)
}

func (r *GrammarRepository) GetByJLPT(ctx context.Context, level int, limit int) ([]*domain.Grammar, error) {
	filter := bson.M{"jlpt": level}

	opt := options.Find().SetLimit(int64(limit))
	cursor, err := r.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, wrapError(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var grammars []*domain.Grammar
	err = cursor.All(ctx, &grammars)
	return grammars, wrapError(err)
}

func (r *GrammarRepository) GetRandom(ctx context.Context, count int) ([]*domain.Grammar, error) {
	filter := bson.M{}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$sample", Value: bson.M{"size": count}}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, wrapError(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var grammars []*domain.Grammar
	err = cursor.All(ctx, &grammars)
	return grammars, wrapError(err)
}

func (r *GrammarRepository) Create(ctx context.Context, g *domain.Grammar) error {
	_, err := r.collection.InsertOne(ctx, g)
	return wrapError(err)
}

func (r *GrammarRepository) BulkCreate(ctx context.Context, grammars []*domain.Grammar) (int64, error) {
	docs := make([]any, len(grammars))
	for i, q := range grammars {
		docs[i] = q
	}

	result, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return 0, wrapError(err)
	}

	return int64(len(result.InsertedIDs)), nil
}

func (r *GrammarRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return wrapError(err)
	}

	if result.DeletedCount == 0 {
		return domain.ErrNotFound
	}

	return nil
}
