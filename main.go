package main

import (
	"fmt"
	"log"
	"time"

	"go-stock-analysis/database"
	"go-stock-analysis/handlers"
	"go-stock-analysis/helpers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	conn, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close()

	fmt.Println("Anslutning till PostgreSQL lyckades!")

	database.InitializeDB(conn)

	// Skapa en ny Gin-router
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
		c.Set("db", conn)
		c.Next()
	})

	// Definiera dina routes
	r.GET("/", helpers.WelcomeMessage)
	r.GET("/stocks", handlers.GetAllStocks)
	r.POST("/addstock", handlers.AddStock)
	r.GET("/getstock", handlers.GetStock)
	r.DELETE("/deletestock", handlers.DeleteStock)
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

	// Log all requests and responses for debugging
	r.Use(func(c *gin.Context) {
		log.Println("Request URL:", c.Request.URL)
		log.Println("Request Method:", c.Request.Method)
		log.Println("Request Headers:", c.Request.Header)
		c.Next()
		log.Println("Response Status:", c.Writer.Status())
		log.Println("Response Headers:", c.Writer.Header())
	})

	// Starta webbservern på port 8085
	r.Run(":8085")
}
