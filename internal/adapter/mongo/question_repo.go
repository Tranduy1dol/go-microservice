package mongo

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuestionRepository struct {
	collection *mongo.Collection
}

func NewQuestionRepository(db *mongo.Database) *QuestionRepository {
	return &QuestionRepository{collection: db.Collection("questions")}
}

func (r *QuestionRepository) GetByID(ctx context.Context, id string) (*domain.Question, error) {
	filter := bson.M{"_id": id}

	var question domain.Question
	err := r.collection.FindOne(ctx, filter).Decode(&question)
	return &question, wrapError(err)
}

func (r *QuestionRepository) GetByJLPT(ctx context.Context, level int, section domain.TestSection, limit int) ([]*domain.Question, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"jlpt": level},
			{"section": section},
		},
	}

	opt := options.Find().SetLimit(int64(limit))
	cursor, err := r.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, wrapError(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var questions []*domain.Question
	err = cursor.All(ctx, &questions)
	return questions, wrapError(err)
}

func (r *QuestionRepository) GetRandom(ctx context.Context, count int) ([]*domain.Question, error) {
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

	var questions []*domain.Question
	err = cursor.All(ctx, &questions)
	return questions, wrapError(err)
}

func (r *QuestionRepository) Create(ctx context.Context, q *domain.Question) error {
	_, err := r.collection.InsertOne(ctx, q)
	return wrapError(err)
}

func (r *QuestionRepository) BulkCreate(ctx context.Context, questions []*domain.Question) (int64, error) {
	docs := make([]any, len(questions))
	for i, q := range questions {
		docs[i] = q
	}

	result, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return 0, wrapError(err)
	}

	return int64(len(result.InsertedIDs)), nil
}

func (r *QuestionRepository) Delete(ctx context.Context, id string) error {
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

func (r *QuestionRepository) List(ctx context.Context, limit, offset int) ([]*domain.Question, int, error) {
	filter := bson.M{}
	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, wrapError(err)
	}

	opts := options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)).SetSort(bson.M{"_id": -1})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, wrapError(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var questions []*domain.Question
	if err := cursor.All(ctx, &questions); err != nil {
		return nil, 0, wrapError(err)
	}

	return questions, int(total), nil
}

func (r *QuestionRepository) Update(ctx context.Context, id string, question *domain.Question) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"type":          question.Type,
		"section":       question.Section,
		"jlpt":          question.JLPT,
		"prompt":        question.Prompt,
		"choices":       question.Choices,
		"correct_index": question.CorrectIndex,
		"explanation":   question.Explanation,
		"tags":          question.Tags,
		"source":        question.Source,
	}}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return wrapError(err)
	}

	if result.MatchedCount == 0 {
		return domain.ErrNotFound
	}

	return nil
}
