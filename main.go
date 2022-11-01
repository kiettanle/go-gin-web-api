package main

import (
	"context"
	"fmt"
	"go-gin-web-api/controllers"
	"go-gin-web-api/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userService    services.UserService
	userController controllers.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	mongoClient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
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

	userCollection = mongoClient.Database("UserManagement").Collection("users")
	userService = services.NewUserService(userCollection, ctx)
	userController = controllers.NewUserController(userService)

	server = gin.Default()
}

func main() {

	defer mongoClient.Disconnect(ctx)

	apiV1 := server.Group("/api/v1")
	userController.RegisterUserRoute(apiV1)

	const PORT = 9090
	log.Fatal(server.Run(":9090"))

	fmt.Println("API server started at PORT: ", PORT)
}
