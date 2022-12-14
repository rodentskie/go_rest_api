package userController

import (
	"fmt"
	middlware "go-rest-api/src/middleware"
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

	newUserId, token, err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "id": &newUserId, "token": token})
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

	fmt.Printf("-----%v\n ", ctx.Request.Header.Get("Request-User-Id")) // user ID from JWT token

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
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

	var id string = ctx.Param("id")

	err := uc.UserService.UpdateUser(&id, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var id string = ctx.Param("id")
	err := uc.UserService.DeleteUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/", uc.CreateUser)
	userroute.GET("/:id", uc.GetSingleUser)
	userroute.GET("/", middlware.ValidateToken(), uc.GetAllUser)
	userroute.PATCH("/:id", uc.UpdateUser)
	userroute.DELETE("/:id", uc.DeleteUser)
}
