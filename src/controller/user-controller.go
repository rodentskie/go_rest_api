package userController

import (
	userModel "go-rest-api/src/models"
	userService "go-rest-api/src/services/users"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	UserService userService.UserService
}

func InitUserController(userservice userService.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user userModel.User
	validate := validator.New()

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := validate.Struct(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := uc.UserService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "create"})
}

func (uc *UserController) GetSingleUser(ctx *gin.Context) {
	var id string = ctx.Param("id")
	user, err := uc.UserService.GetSingleUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUser(ctx *gin.Context) {
	users, err := uc.UserService.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/", uc.CreateUser)
	userroute.GET("/:id", uc.GetSingleUser)
	userroute.GET("/", uc.GetAllUser)
	userroute.PATCH("/:id", uc.UpdateUser)
	userroute.DELETE("/:id", uc.DeleteUser)
}
