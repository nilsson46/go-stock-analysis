package handlers

import (
	"net/http"

	"go-stock-analysis/database"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetStocks(c *gin.Context) {
	conn := c.MustGet("db").(*pgxpool.Pool)
	stockList, err := database.GetStocksFromDB(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, stockList)
}

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

	exists, err := database.StockExists(conn, stock.Name, stock.Symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Stock already exists"})
		return
	}

	err = database.AddStock(conn, stock.Name, stock.Price, stock.Symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock added successfully"})
}
