package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// PostgreSQL driver

	"go-stock-analysis/helpers"
	"go-stock-analysis/stocks"
)

func main() {

	conn, err := ConnectDB()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Anslutning till PostgreSQL lyckades!")

	// Create table if not exists
	_, err = conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS stocks (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		price DECIMAL,
		symbol VARCHAR(10)
	)`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}

	// Insert a stock entry
	_, err = conn.Exec(context.Background(), `INSERT INTO stocks (name, price, symbol) VALUES ($1, $2, $3)`, "Example Stock", 100.50, "EXMPL")
	if err != nil {
		log.Fatalf("Unable to insert stock: %v\n", err)
	}

	// Retrieve and print stock entries
	rows, err := conn.Query(context.Background(), `SELECT name, price, symbol FROM stocks`)
	if err != nil {
		log.Fatalf("Unable to retrieve stocks: %v\n", err)
	}
	defer rows.Close()

	fmt.Println("Stocks in database:")
	for rows.Next() {
		var name string
		var price float64
		var symbol string
		err := rows.Scan(&name, &price, &symbol)
		if err != nil {
			log.Fatalf("Unable to scan row: %v\n", err)
		}
		fmt.Printf("Name: %s, Price: %.2f, Symbol: %s\n", name, price, symbol)
	}

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		log.Fatalf("Error encountered during rows iteration: %v\n", err)
	}

	// Skapa en ny Gin-router
	r := gin.Default()

	// Hantera en GET-förfrågan på roten ("/")
	r.GET("/", helpers.WelcomeMessage)

	// Hantera en GET-förfrågan på "/stocks"
	r.GET("/stocks", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, stocks.GetStocks())
	})

	// Starta webbservern på port 8085
	r.Run(":8085")
}
