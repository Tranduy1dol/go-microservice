package mongo

import (
	"context"
	"time"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SRSRepository struct {
	collection *mongo.Collection
}

func NewSRSRepository(db *mongo.Database) *SRSRepository {
	return &SRSRepository{collection: db.Collection("srs_cards")}
}

func (r *SRSRepository) Save(ctx context.Context, card *domain.SRSCard) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": card.ID}
	update := bson.M{"$set": card}

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return wrapError(err)
}

func (r *SRSRepository) GetDueCards(ctx context.Context, userID string, limit int) ([]*domain.SRSCard, error) {
	filter := bson.M{"user_id": userID, "due_date": bson.M{"$lte": time.Now()}}
	opts := options.Find().SetSort(bson.M{"due_date": 1}).SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, wrapError(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var cards []*domain.SRSCard
	err = cursor.All(ctx, &cards)
	return cards, wrapError(err)
}

func (r *SRSRepository) GetByIDAndUser(ctx context.Context, id, userID string) (*domain.SRSCard, error) {
	filter := bson.M{"_id": id, "user_id": userID}

	var card domain.SRSCard
	err := r.collection.FindOne(ctx, filter).Decode(&card)
	return &card, wrapError(err)
}

func (r *SRSRepository) GetByWordAndUser(ctx context.Context, wordID, userID string) (*domain.SRSCard, error) {
	filter := bson.M{"word_id": wordID, "user_id": userID}

	var card domain.SRSCard
	err := r.collection.FindOne(ctx, filter).Decode(&card)
	return &card, wrapError(err)
}

func (r *SRSRepository) GetDueCardsCount(ctx context.Context, userID string) (int64, error) {
	filter := bson.M{
		"user_id":  userID,
		"due_date": bson.M{"$lte": time.Now()},
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	return count, wrapError(err)
}

func (r *SRSRepository) GetUserWordIDs(ctx context.Context, userID string) ([]string, error) {
	filter := bson.M{"user_id": userID}
	cursor, err := r.collection.Find(ctx, filter, options.Find().SetProjection(bson.M{"word_id": 1}))
	if err != nil {
		return nil, wrapError(err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var results []struct {
		WordID string `bson:"word_id"`
	}
	if err := cursor.All(ctx, &results); err != nil {
		return nil, wrapError(err)
	}

	ids := make([]string, len(results))
	for i, r := range results {
		ids[i] = r.WordID
	}

	return ids, nil
}
