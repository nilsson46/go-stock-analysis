package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeMessage(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the stock information!")
}
