package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	database := Connect()

	userCollection := database.Collection("users")
	todoCollection := database.Collection("todos")

	createIndexForUserCollection(userCollection)
	createIndexForTodoCollection(todoCollection)
}

func createIndexForUserCollection(collection *mongo.Collection) {
	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func createIndexForTodoCollection(collection *mongo.Collection) {
	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
