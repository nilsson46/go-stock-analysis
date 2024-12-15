package helpers

import (
	"github.com/gin-gonic/gin"
)

func WelcomeMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the Stock Analysis API",
	})
}
