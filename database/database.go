package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// ConnectDB ansluter till PostgreSQL-databasen
func ConnectDB() (*pgxpool.Pool, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatalf("DATABASE_URL is not set")
	}
	conn, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// InitializeDB initialiserar databasen
func InitializeDB(conn *pgxpool.Pool) {
	// Lägg till eventuell initialisering av databasen här
}

// StockExists kontrollerar om en aktie med samma namn eller symbol redan finns i databasen
func StockExists(conn *pgxpool.Pool, name, symbol string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM stocks WHERE name=$1 OR symbol=$2)`
	err := conn.QueryRow(context.Background(), query, name, symbol).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if stock exists: %v", err)
	}
	return exists, err
}

// AddStock lägger till en ny aktie i databasen
func AddStock(conn *pgxpool.Pool, name string, price float64, symbol string) error {
	query := `INSERT INTO stocks (name, price, symbol) VALUES ($1, $2, $3)`
	_, err := conn.Exec(context.Background(), query, name, price, symbol)
	if err != nil {
		log.Printf("Error adding stock: %v", err)
	}
	return err
}

// GetStock hämtar en specifik aktie baserat på dess namn eller symbol
func GetStock(conn *pgxpool.Pool, name, symbol string) (map[string]interface{}, error) {
	var stock map[string]interface{}
	query := `SELECT name, price, symbol FROM stocks WHERE name=$1 OR symbol=$2`
	row := conn.QueryRow(context.Background(), query, name, symbol)
	var stockName string
	var stockPrice float64
	var stockSymbol string
	err := row.Scan(&stockName, &stockPrice, &stockSymbol)
	if err != nil {
		return nil, err
	}
	stock = map[string]interface{}{
		"name":   stockName,
		"price":  stockPrice,
		"symbol": stockSymbol,
	}
	return stock, nil
}

// GetStocksFromDB hämtar alla aktier från databasen
func GetStocksFromDB(conn *pgxpool.Pool) ([]map[string]interface{}, error) {
	rows, err := conn.Query(context.Background(), "SELECT name, price, symbol FROM stocks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []map[string]interface{}
	for rows.Next() {
		var name string
		var price float64
		var symbol string
		err := rows.Scan(&name, &price, &symbol)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, map[string]interface{}{
			"name":   name,
			"price":  price,
			"symbol": symbol,
		})
	}
	return stocks, nil
}

// DeleteStock tar bort en aktie från databasen
func DeleteStock(conn *pgxpool.Pool, name, symbol string) error {
	query := `DELETE FROM stocks WHERE name=$1 OR symbol=$2`
	_, err := conn.Exec(context.Background(), query, name, symbol)
	if err != nil {
		log.Printf("Error deleting stock: %v", err)
	}
	return err
}
