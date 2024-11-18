package handlers

import (
	"go-stock-analysis/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

// AddStock lägger till en ny aktie
func AddStock(c *gin.Context) {
	conn := c.MustGet("db").(*pgxpool.Pool)
	var stock struct {
		Name   string  `json:"name"`
		Price  float64 `json:"price"`
		Symbol string  `json:"symbol"`
	}

	if err := c.BindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Kontrollera att alla parametrar är fyllda
	if stock.Name == "" || stock.Symbol == "" || stock.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Kontrollera om namnet eller symbolen redan finns i databasen
	exists, err := database.StockExists(conn, stock.Name, stock.Symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Stock with the same name or symbol already exists"})
		return
	}

	err = database.AddStock(conn, stock.Name, stock.Price, stock.Symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock added successfully"})
}

// GetStock hämtar en specifik aktie baserat på dess namn eller symbol
func GetStock(c *gin.Context) {
	conn := c.MustGet("db").(*pgxpool.Pool)
	name := c.Query("name")
	symbol := c.Query("symbol")

	if name == "" && symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name or symbol is required"})
		return
	}

	stock, err := database.GetStock(conn, name, symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if stock == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, stock)
}

// GetAllStocks hämtar alla aktier från databasen
func GetAllStocks(c *gin.Context) {
	conn := c.MustGet("db").(*pgxpool.Pool)
	stocks, err := database.GetStocksFromDB(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stocks)
}

// DeleteStock tar bort en aktie från databasen
func DeleteStock(c *gin.Context) {
	conn := c.MustGet("db").(*pgxpool.Pool)
	name := c.Query("name")
	symbol := c.Query("symbol")

	if name == "" && symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name or symbol is required"})
		return
	}

	err := database.DeleteStock(conn, name, symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
