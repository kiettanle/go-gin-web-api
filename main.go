package main

import (
	"context"
	"fmt"
	"go-gin-web-api/config"
	"go-gin-web-api/controllers"
	"go-gin-web-api/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
	mongoClient, _ := config.GetMongoClient()
	userCollection = mongoClient.Database("UserManagement").Collection("users")
	userService = services.NewUserService(userCollection, ctx)
	userController = controllers.NewUserController(userService)

	server = gin.Default()
}

func main() {

	defer mongoClient.Disconnect(ctx)

	apiV1 := server.Group("/api/v1")
	userController.RegisterUserRoute(apiV1)

	// TODO: Update to get PORT from .env
	const PORT = 9090
	log.Fatal(server.Run(":9090"))

	fmt.Println("API server started at PORT: ", PORT)
}
