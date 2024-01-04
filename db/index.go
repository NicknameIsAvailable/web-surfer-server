package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func ConnectDB(uri string) error {
	var err error

	clientOptions := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	return err
}

func GetClient() *mongo.Client {
	return client
}
