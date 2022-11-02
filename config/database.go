package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx         context.Context
	mongoClient *mongo.Client
	err         error
)

// Upper case the first letter to expose func
func GetMongoClient() (*mongo.Client, error) {
	ctx = context.TODO()
	// TODO: Update to get url from .env
	connection := options.Client().ApplyURI("mongodb://admin:Init123456@localhost:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false")
	mongoClient, err = mongo.Connect(ctx, connection)

	if err != nil {
		log.Fatal("Cannot connect to mongo db")
	}

	err = mongoClient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("Cannot ping to mongo db")
	}

	fmt.Println("Mongo db is ready to connect!!!")

	return mongoClient, err
}
