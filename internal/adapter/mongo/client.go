package mongo

import (
	"context"
	"time"

	"github.com/Tranduy1dol/kotoba-press-core/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var log = logger.New(logger.ComponentMongo)

func NewClient(ctx context.Context, uri, dbName string) (*mongo.Client, *mongo.Database, error) {
	clientOpts := options.Client().ApplyURI(uri).SetConnectTimeout(10 * time.Second).SetServerSelectionTimeout(5 * time.Second)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Error("failed to ping mongodb", "error", err)
		return nil, nil, err
	}

	log.Info("connected to mongodb", "database", dbName)
	return client, client.Database(dbName), nil
}

func Close(client *mongo.Client) error {
	if client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := client.Disconnect(ctx); err != nil {
		log.Error("failed to disconnect from mongodb", "error", err)
		return err
	}
	
	log.Info("disconnected from mongodb")
	return nil
}
