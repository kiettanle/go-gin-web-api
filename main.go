package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"

	"go-gin-web-api/config"
	"go-gin-web-api/controllers"
	"go-gin-web-api/docs"
	"go-gin-web-api/services"
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

// @title           Go Gin Web API
// @version         1.0
// @description     This is a sample server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9090
// @BasePath  /api/v1
func main() {

	defer mongoClient.Disconnect(ctx)

	// Health check
	server.GET("api/v1/health-check", controllers.HealthCheck)

	apiV1 := server.Group("/api/v1")
	userController.RegisterUserRoute(apiV1)

	docs.SwaggerInfo.BasePath = "/api/v1"

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":9090")

	// TODO: Update to get PORT from .env
	const PORT = 9090
	log.Fatal(server.Run(":9090"))

	fmt.Println("API server started at PORT: ", PORT)
}
