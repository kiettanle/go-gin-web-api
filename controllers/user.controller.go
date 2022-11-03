package controllers

import (
	"go-gin-web-api/models"
	"go-gin-web-api/services"
	"go-gin-web-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) Create(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createdUser, err := uc.UserService.Create(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
	return
}

func (uc *UserController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := uc.UserService.Get(&id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	updatedUser, err := uc.UserService.Update(&id, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (uc *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := uc.UserService.Delete(&id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User has been deleted successfully."})
}

func (uc *UserController) CreateBulkConcurrency(ctx *gin.Context) {
	defer utils.TimeTrack(time.Now(), "CreateBulkConcurrency")
	var users []models.User

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for i := 0; i < len(users); i++ {
		go uc.UserService.Create(&users[i])
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Users have been created successfully."})
}

func (uc *UserController) CreateBulkWithoutConcurrency(ctx *gin.Context) {
	defer utils.TimeTrack(time.Now(), "CreateBulkWithoutConcurrency")
	var users []models.User

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for i := 0; i < len(users); i++ {
		uc.UserService.Create(&users[i])
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Users have been created successfully."})
}

func (uc *UserController) RegisterUserRoute(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")
	userRoute.POST("/", uc.Create)
	userRoute.POST("/bulk/concurrency", uc.CreateBulkConcurrency)
	userRoute.POST("/bulk/without-concurrency", uc.CreateBulkWithoutConcurrency)
	userRoute.GET("/", uc.GetAll)
	userRoute.GET("/:id", uc.Get)
	userRoute.PUT("/:id", uc.Update)
	userRoute.DELETE("/:id", uc.Delete)
}
