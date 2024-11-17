package routes

import (
	"go-stock-analysis/handlers"
	"go-stock-analysis/helpers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	router.GET("/", helpers.WelcomeMessage)

	router.GET("/stocks", handlers.GetStock)

	router.POST("/addstock", handlers.AddStock)
}
