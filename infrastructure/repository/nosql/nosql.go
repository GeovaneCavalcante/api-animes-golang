package nosql

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NoSqlRepo nosql repo
type NoSqlRepo struct {
	db *mongo.Database
}

// NewNoSqlRepo create a new nosqlrepo
func NewNoSqlRepo(db *mongo.Database) *NoSqlRepo {
	return &NoSqlRepo{
		db: db,
	}
}

// ConnectMongoDB connect to database
func ConnectMongoDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(""))

	if err != nil {
		return nil, err
	}
	db := client.Database("animedb")

	return db, err
}
