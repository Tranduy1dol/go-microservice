package mongo

import (
	"context"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ParagraphRepository struct {
	collection *mongo.Collection
}

func NewParagraphRepository(db *mongo.Database) *ParagraphRepository {
	return &ParagraphRepository{collection: db.Collection("paragraph")}
}

func (r *ParagraphRepository) GetByID(ctx context.Context, id string) (*domain.Paragraph, error) {
	filter := bson.M{"_id": id}

	var paragraph domain.Paragraph
	err := r.collection.FindOne(ctx, filter).Decode(&paragraph)
	if err != nil {
		return nil, err
	}

	return &paragraph, nil
}

func (r *ParagraphRepository) GetByJLPT(ctx context.Context, level int, limit, offset int) ([]*domain.Paragraph, error) {
	filter := bson.M{"jlpt": level}

	opt := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := r.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	var paragraphs []*domain.Paragraph
	if err := cursor.All(ctx, &paragraphs); err != nil {
		return nil, err
	}

	return paragraphs, nil
}

func (r *ParagraphRepository) GetRandom(ctx context.Context, count int) ([]*domain.Paragraph, error) {
	filter := bson.M{}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$sample", Value: bson.M{"size": count}}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	var paragraph []*domain.Paragraph
	if err := cursor.All(ctx, &paragraph); err != nil {
		return nil, err
	}

	return paragraph, nil
}

func (r *ParagraphRepository) Create(ctx context.Context, p *domain.Paragraph) error {
	_, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (r *ParagraphRepository) BulkCreate(ctx context.Context, paragraphs []*domain.Paragraph) (int64, error) {
	docs := make([]any, len(paragraphs))
	for i, q := range paragraphs {
		docs[i] = q
	}

	result, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return 0, err
	}

	return int64(len(result.InsertedIDs)), nil
}

func (r *ParagraphRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
