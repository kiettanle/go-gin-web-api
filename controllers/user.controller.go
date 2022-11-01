package controllers

import (
	"fmt"
	"go-gin-web-api/models"
	"go-gin-web-api/services"
	"net/http"

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

	fmt.Println(user)

	err := uc.UserService.Create(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &user)
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
	ctx.JSON(200, "nil")
}

func (uc *UserController) Update(ctx *gin.Context) {
	ctx.JSON(200, "nil")
}

func (uc *UserController) Delete(ctx *gin.Context) {
	ctx.JSON(200, "nil")
}

func (uc *UserController) RegisterUserRoute(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")
	userRoute.POST("/", uc.Create)
	userRoute.GET("/", uc.GetAll)
	userRoute.GET("/:id", uc.Get)
	userRoute.PUT("/", uc.Update)
	userRoute.DELETE("/", uc.Delete)
}
