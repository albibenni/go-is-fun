// Package config provides database connection and collection management for MongoDB.
package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB holds the MongoDB client and database configuration
type DB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// InitDB establishes a connection to MongoDB and returns a DB instance.
// It uses the MONGODB_URI environment variable or defaults to localhost:27017.
// Connection includes proper pooling configuration and error handling.
//
// Returns:
//   - *DB: Database instance with client and database references
//   - error: Any connection error encountered
func InitDB() (*DB, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("MONGODB_DATABASE")
	if dbName == "" {
		dbName = "usersdb"
	}

	log.Printf("Attempting to connect to MongoDB at %s...", uri)

	clientOptions := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(20).
		SetMinPoolSize(5).
		SetMaxConnIdleTime(30 * time.Second).
		SetConnectTimeout(10 * time.Second).
		SetServerSelectionTimeout(5 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to verify connection
	if err = client.Ping(ctx, nil); err != nil {
		client.Disconnect(ctx)
		return nil, fmt.Errorf("MongoDB is not reachable: %w", err)
	}

	log.Println("Successfully connected to MongoDB!")

	return &DB{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

// Close gracefully closes the MongoDB connection.
//
// Parameters:
//   - ctx: Context for the disconnect operation
//
// Returns:
//   - error: Any error encountered during disconnect
func (db *DB) Close(ctx context.Context) error {
	if db.Client == nil {
		return nil
	}

	log.Println("Closing MongoDB connection...")
	if err := db.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}

	log.Println("MongoDB connection closed successfully")
	return nil
}

// GetCollection returns a MongoDB collection instance for the specified collection name.
//
// Parameters:
//   - collectionName: The name of the collection to retrieve
//
// Returns:
//   - *mongo.Collection: A pointer to the MongoDB collection
//   - error: Error if database is not initialized
func (db *DB) GetCollection(collectionName string) (*mongo.Collection, error) {
	if db.Database == nil {
		return nil, fmt.Errorf("database is not initialized")
	}
	return db.Database.Collection(collectionName), nil
}

func GetUserCollection(db *DB) *mongo.Collection {

	userCollection, err := db.GetCollection("users")
	if err != nil {
		log.Fatal("Error with users collection", err)
	}
	return userCollection
}
