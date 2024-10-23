package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"

	"go-stock-analysis/database"
	"go-stock-analysis/helpers"
)

type StockSearchResult struct {
	TotalNumberOfHits int `json:"totalNumberOfHits"`
	Hits              []struct {
		Type       string `json:"type"`
		Instrument struct {
			Identifier string `json:"identifier"`
			Name       string `json:"name"`
			Price      struct {
				LastPrice float64 `json:"lastPrice"`
			} `json:"price"`
			Symbol string `json:"symbol"`
		} `json:"instrument"`
	} `json:"hits"`
}

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

	// Hantera en GET-förfrågan på roten ("/")
	r.GET("/", helpers.WelcomeMessage)

	// Hantera en GET-förfrågan på "/stocks"
	r.GET("/stocks", func(c *gin.Context) {
		stockList, err := database.GetStocksFromDB(conn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, stockList)
	})

	// Hantera en POST-förfrågan på "/addstock"
	r.POST("/addstock", func(c *gin.Context) {
		var stock struct {
			Name   string  `json:"name"`
			Price  float64 `json:"price"`
			Symbol string  `json:"symbol"`
		}

		if err := c.BindJSON(&stock); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		err := database.AddStock(conn, stock.Name, stock.Price, stock.Symbol)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Stock added successfully"})
	})

	// Hantera en GET-förfrågan på "/search"
	r.GET("/search", func(c *gin.Context) {
		stockName := c.Query("name")
		if stockName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Stock name is required"})
			return
		}

		client := resty.New()
		apiURL := fmt.Sprintf("https://www.avanza.se/_mobile/market/search/%s", stockName)

		resp, err := client.R().Get(apiURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from API"})
			return
		}

		var result StockSearchResult
		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse API response"})
			return
		}

		c.IndentedJSON(http.StatusOK, result)
	})

	// Starta webbservern på port 8085
	r.Run(":8085")
}
