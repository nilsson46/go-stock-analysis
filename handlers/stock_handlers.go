package handlers

import (
	"go-stock-analysis/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddStock adds a new stock
func AddStock(c *gin.Context) {
	db := c.MustGet("db").(database.DB)
	var stock struct {
		Name   string  `json:"name"`
		Price  float64 `json:"price"`
		Symbol string  `json:"symbol"`
	}

	if err := c.BindJSON(&stock); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if stock.Name == "" || stock.Symbol == "" || stock.Price == 0 {
		log.Println("Missing fields in request payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	exists, err := database.StockExists(db, stock.Name, stock.Symbol)
	if err != nil {
		log.Println("Error checking if stock exists:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		log.Println("Stock already exists")
		c.JSON(http.StatusConflict, gin.H{"error": "Stock with the same name or symbol already exists"})
		return
	}

	err = database.AddStock(db, stock.Name, stock.Price, stock.Symbol)
	if err != nil {
		log.Println("Error adding stock:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Stock added successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Stock added successfully"})
}

// GetStock retrieves a specific stock by its name or symbol
func GetStock(c *gin.Context) {
	db := c.MustGet("db").(database.DB)
	name := c.Query("name")
	symbol := c.Query("symbol")

	if name == "" && symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name or symbol is required"})
		return
	}

	stock, err := database.GetStock(db, name, symbol)
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

// GetAllStocks retrieves all stocks from the database
func GetAllStocks(c *gin.Context) {
	db := c.MustGet("db").(database.DB)
	stocks, err := database.GetStocksFromDB(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stocks)
}

// DeleteStock deletes a stock by its symbol
func DeleteStock(c *gin.Context) {
	db := c.MustGet("db").(database.DB)
	symbol := c.Query("symbol")

	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Symbol is required"})
		return
	}

	err := database.DeleteStockBySymbol(db, symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock deleted successfully"})
}
