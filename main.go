package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-stock-analysis/helpers"

	"go-stock-analysis/stocks"
)

func main() {
	// Skapa en ny Gin-router
	r := gin.Default()

	// Hantera en GET-förfrågan på roten ("/")
	r.GET("/", helpers.WelcomeMessage)

	// Hantera en GET-förfrågan på "/stocks"
	r.GET("/stocks", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, stocks.GetStocks())
	})

	// Starta webbservern på port 8080
	r.Run(":8086")
}
