package core

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type MongoConfig struct {
	Database string
	MongoURI string
}

func (db *MongoConfig) Configure(dbName string, mongoURI string) *MongoConfig {
	db.Database = dbName
	db.MongoURI = mongoURI

	return db
}

func (db *MongoConfig) Connect() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.MongoURI))

	if err != nil {
		log.Fatalf("An error occurred while connecting to database: \n %v", err)
	}

	return client.Database(db.Database)
}
