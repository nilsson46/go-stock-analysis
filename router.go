package main

import (
	"go-stock-analysis/database"
	"go-stock-analysis/handlers"
	"go-stock-analysis/helpers"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter(db database.DB) *gin.Engine {
	r := gin.Default()

	// Apply CORS middleware before defining routes
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Middleware för att lägga till databaskopplingen i context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Definiera dina routes
	r.GET("/", helpers.WelcomeMessage)
	r.GET("/stocks", handlers.GetAllStocks)
	r.POST("/addstock", handlers.AddStock)
	r.GET("/getstock", handlers.GetStock)
	r.DELETE("/deletestock", handlers.DeleteStock)
	r.DELETE("/deletestock/:symbol", handlers.DeleteStock)
	r.PUT("/updatestockprice", handlers.UpdateStockPrice)
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/debug", func(c *gin.Context) {
		log.Println("Debug route called")
		log.Println("Request headers:", c.Request.Header)
		c.JSON(200, gin.H{
			"message": "Debug successful",
			"headers": c.Request.Header,
		})
	})

	return r
}
