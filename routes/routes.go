package main

import (
	"go-stock-analysis/database"
	"go-stock-analysis/handlers"
	"go-stock-analysis/helpers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter(conn *database.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	})

	r.GET("/", helpers.WelcomeMessage)
	r.GET("/stocks", handlers.GetAllStocks)
	r.POST("/addstock", handlers.AddStock)
	r.GET("/getstock", handlers.GetStock)
	r.DELETE("/deletestock", handlers.DeleteStock)

	return r
}
