package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "delete",
	})
}
