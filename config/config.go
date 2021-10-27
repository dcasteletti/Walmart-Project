package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	host     = "localhost"
	port     = 27017
	database = "local"
)

func GetCollection(collection string) *mongo.Collection {
	url := fmt.Sprintf("mongodb://%s:%d", host, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil
	}

	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(context)
	if err != nil {
		return nil
	}

	return client.Database(database).Collection(collection)
}
