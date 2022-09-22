package server

import (
	userController "go-rest-api/src/controller/users"
	db "go-rest-api/src/data"

	"github.com/gin-gonic/gin"
)

func Server() {
	db.Start()

	router := gin.Default()

	// User Routes and Enpoints
	router.GET("/users", userController.GetAllUsers)
	router.GET("/users/:id", userController.GetSingleUser)
	router.POST("/users", userController.CreateUser)
	router.PATCH("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	router.Run(":8080")
}
