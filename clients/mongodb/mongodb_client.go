package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	dbName = "items"
	dbCollection = "items"
)

var (
	collection *mongo.Collection
	ctx = context.TODO()
)

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	collection = client.Database(dbName).Collection(dbCollection)
}

func InsertDocument(document interface{}) (string, error) {
	result, err := collection.InsertOne(ctx, document)
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}
