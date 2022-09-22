package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "create",
	})
}
