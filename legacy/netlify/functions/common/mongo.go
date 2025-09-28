package common

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	mu          sync.Mutex
)

func GetMongoClient(ctx context.Context) (*mongo.Client, error) {
	mu.Lock()
	defer mu.Unlock()

	if mongoClient != nil {
		return mongoClient, nil
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, fmt.Errorf("MONGODB_URI is empty")
	}

	clientOpts := options.Client().ApplyURI(uri)
	ctx2, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx2, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}

	if err := c.Ping(ctx2, nil); err != nil {
		_ = c.Disconnect(ctx2)
		return nil, fmt.Errorf("mongo ping: %w", err)
	}

	mongoClient = c
	return mongoClient, nil
}

func DisconnectClient(ctx context.Context) error {
	mu.Lock()
	defer mu.Unlock()
	if mongoClient == nil {
		return nil
	}
	err := mongoClient.Disconnect(ctx)
	mongoClient = nil
	return err
}
