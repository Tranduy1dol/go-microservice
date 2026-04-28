package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewClient(ctx context.Context, uri, dbName string) (*mongo.Database, error) {
}

func Close(client *mongo.Client) error {
}
