package mongo

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
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
	return &paragraph, wrapError(err)
}

func (r *ParagraphRepository) GetByJLPT(ctx context.Context, level int, limit, offset int) ([]*domain.Paragraph, error) {
	filter := bson.M{"jlpt": level}

	opt := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := r.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, wrapError(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var paragraphs []*domain.Paragraph
	err = cursor.All(ctx, &paragraphs)
	return paragraphs, wrapError(err)
}

func (r *ParagraphRepository) GetRandom(ctx context.Context, count int) ([]*domain.Paragraph, error) {
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

	var paragraph []*domain.Paragraph
	err = cursor.All(ctx, &paragraph)
	return paragraph, wrapError(err)
}

func (r *ParagraphRepository) Create(ctx context.Context, p *domain.Paragraph) error {
	_, err := r.collection.InsertOne(ctx, p)
	return wrapError(err)
}

func (r *ParagraphRepository) BulkCreate(ctx context.Context, paragraphs []*domain.Paragraph) (int64, error) {
	docs := make([]any, len(paragraphs))
	for i, q := range paragraphs {
		docs[i] = q
	}

	result, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return 0, wrapError(err)
	}

	return int64(len(result.InsertedIDs)), nil
}

func (r *ParagraphRepository) Delete(ctx context.Context, id string) error {
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

func (r *ParagraphRepository) List(ctx context.Context, limit, offset int) ([]*domain.Paragraph, int, error) {
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

	var paragraphs []*domain.Paragraph
	if err := cursor.All(ctx, &paragraphs); err != nil {
		return nil, 0, wrapError(err)
	}

	return paragraphs, int(total), nil
}

func (r *ParagraphRepository) Update(ctx context.Context, id string, paragraph *domain.Paragraph) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"title":     paragraph.Title,
		"content":   paragraph.Content,
		"jlpt":      paragraph.JLPT,
		"questions": paragraph.Questions,
		"tags":      paragraph.Tags,
		"source":    paragraph.Source,
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
