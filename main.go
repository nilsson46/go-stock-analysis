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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
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

	// Starta webbservern på port 8085
	r.Run(":8085")
}
