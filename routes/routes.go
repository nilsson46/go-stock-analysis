package routes

import (
	"go-stock-analysis/handlers"
	"go-stock-analysis/helpers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	router.GET("/", helpers.WelcomeMessage)

	router.GET("/stocks", handlers.GetAllStocks)

	router.POST("/addstock", handlers.AddStock)

	router.DELETE("/deletestock", handlers.DeleteStock)
}
