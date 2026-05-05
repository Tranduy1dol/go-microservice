package mongo

import (
	"context"
	"time"

	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	filter := bson.M{"_id": id}

	var user domain.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, wrapError(err)
	}

	return &user, nil
}

func (r *UserRepository) GetByGoogleID(ctx context.Context, googleID string) (*domain.User, error) {
	filter := bson.M{"google_id": googleID}

	var user domain.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, wrapError(err)
	}

	return &user, nil
}

func (r *UserRepository) Upsert(ctx context.Context, googleID, email, name, pictureURL string) (*domain.User, error) {
	filter := bson.M{"google_id": googleID}

	update := bson.M{
		"$set": bson.M{
			"email":       email,
			"name":        name,
			"picture_url": pictureURL,
			"updated_at":  time.Now(),
		},
		"$setOnInsert": bson.M{
			"_id":        googleID, // Use GoogleID as the _id since it is unique
			"google_id":  googleID,
			"role":       "user",
			"created_at": time.Now(),
			"study_progress": domain.Progress{
				JLPTLevel:      5,
				WeeklyMinutes:  []int{0, 0, 0, 0, 0, 0, 0},
				RecentActivity: []domain.ActiviyLog{},
			},
		},
	}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var user domain.User
	err := r.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&user)
	return &user, wrapError(err)
}

func (r *UserRepository) UpdateProgress(ctx context.Context, userID string, progress domain.Progress) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"study_progress": progress}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return wrapError(err)
}
