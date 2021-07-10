package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

const (
	dbName       = "items"
	dbCollection = "items"
)

var (
	dbHost       = os.Getenv("DB_HOST")
	ctx          = context.TODO()
	collection *mongo.Collection
)

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://" + dbHost)
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

func GetDocument(id string) (*mongo.SingleResult, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := collection.FindOne(ctx, bson.M{"_id": objectId})
	return result, nil
}
