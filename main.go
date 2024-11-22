package main

import (
	"fmt"
	"log"

	"go-stock-analysis/database"
	"go-stock-analysis/handlers"
	"go-stock-analysis/helpers"

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

	// Starta webbservern på port 8085
	r.Run(":8085")
}
