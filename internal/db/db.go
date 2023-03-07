package db

import (
	"auth-service/internal/config"

	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewDBConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetEnvVar("MONGOURI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

func GetCollection(c *mongo.Client, collectionName string) *mongo.Collection {
	return c.Database(config.GetEnvVar("MONGO_DB_NAME")).Collection(collectionName)
}
