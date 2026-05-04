package mongo

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	return &word, wrapError(err)
}

func (r *WordRepository) GetRandom(ctx context.Context, count int) ([]*domain.Word, error) {
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

	var words []*domain.Word
	err = cursor.All(ctx, &words)
	return words, wrapError(err)
}

func (r *WordRepository) GetByKanji(ctx context.Context, kanji string) (*domain.Word, error) {
	filter := bson.M{"kanji.text": kanji}

	var word domain.Word
	err := r.collection.FindOne(ctx, filter).Decode(&word)
	return &word, wrapError(err)
}

func (r *WordRepository) GetByReading(ctx context.Context, reading string) ([]*domain.Word, error) {
	filter := bson.M{"readings.text": reading}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, wrapError(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var words []*domain.Word
	err = cursor.All(ctx, &words)
	return words, wrapError(err)
}

func (r *WordRepository) Search(ctx context.Context, query string, limit int) ([]*domain.Word, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"kanji.text": bson.M{"$regex": query, "$options": "i"}},
			{"readings.text": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	opts := options.Find().SetLimit(int64(limit))
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, wrapError(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var words []*domain.Word
	err = cursor.All(ctx, &words)
	return words, wrapError(err)
}

func (r *WordRepository) GetByJLPT(ctx context.Context, level int, limit, offset int) ([]*domain.Word, int, error) {
	filter := bson.M{"jlpt": level}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, wrapError(err)
	}

	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, wrapError(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var words []*domain.Word
	err = cursor.All(ctx, &words)
	return words, int(total), wrapError(err)
}

func (r *WordRepository) Create(ctx context.Context, word *domain.Word) error {
	_, err := r.collection.InsertOne(ctx, word)
	return wrapError(err)
}

func (r *WordRepository) BulkCreate(ctx context.Context, words []*domain.Word) (int64, error) {
	docs := make([]any, len(words))
	for i, w := range words {
		docs[i] = w
	}

	result, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return 0, wrapError(err)
	}

	return int64(len(result.InsertedIDs)), nil
}

func (r *WordRepository) Delete(ctx context.Context, id string) error {
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
