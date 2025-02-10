package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

func ConnectToDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		return err
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	Client = client
	DB = client.Database("hots")

	log.Println("Successfully connected to MongoDB")
	return nil
}

// DisconnectDB closes the database connection
func DisconnectDB() error {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := Client.Disconnect(ctx); err != nil {
			return err
		}
	}
	return nil
}

// GetCollection returns a handle to a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
