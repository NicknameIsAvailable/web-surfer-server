package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDocument(collectionName string, filter bson.D) (*mongo.SingleResult, error) {
	client := GetClient()
	if client == nil {
		return nil, fmt.Errorf("database client is nil")
	}

	collection := client.Database("mydatabase").Collection(collectionName)
	result := collection.FindOne(context.Background(), filter)
	return result, nil
}
