package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// ConnectDB establishes a connection to the PostgreSQL database using pgx.
func ConnectDB() (*pgxpool.Pool, error) {
	connStr := "postgres://postgres:1234@localhost:5433/stock_analysis_db?sslmode=disable"
	return pgxpool.Connect(context.Background(), connStr)
}

// InitializeDB initializes the database by creating the necessary tables and updating the schema if needed.
func InitializeDB(conn *pgxpool.Pool) {
	_, err := conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS stocks (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50),
        price DECIMAL,
        symbol VARCHAR(10), 
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        direktavkastning DECIMAL,
        ep DECIMAL,
        egetkapital DECIMAL
    )`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}

	_, err = conn.Exec(context.Background(), `ALTER TABLE stocks
        ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        ADD COLUMN IF NOT EXISTS direktavkastning DECIMAL,
        ADD COLUMN IF NOT EXISTS ep DECIMAL,
        ADD COLUMN IF NOT EXISTS egetkapital DECIMAL`)
	if err != nil {
		log.Fatalf("Unable to update table schema: %v\n", err)
	}

	_, err = conn.Exec(context.Background(), `INSERT INTO stocks (name, price, symbol) VALUES ($1, $2, $3)`, "Example Stock", 100.50, "EXMPL")
	if err != nil {
		log.Fatalf("Unable to insert stock: %v\n", err)
	}
}

// GetStocksFromDB retrieves all stocks from the database.
func GetStocksFromDB(conn *pgxpool.Pool) ([]map[string]interface{}, error) {
	rows, err := conn.Query(context.Background(), `SELECT name, price, symbol FROM stocks`)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve stocks: %v", err)
	}
	defer rows.Close()

	var stocks []map[string]interface{}
	for rows.Next() {
		var name string
		var price float64
		var symbol string
		err := rows.Scan(&name, &price, &symbol)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		stocks = append(stocks, map[string]interface{}{
			"name":   name,
			"price":  price,
			"symbol": symbol,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during rows iteration: %v", err)
	}

	return stocks, nil
}

// AddStock adds a new stock to the database.
func AddStock(conn *pgxpool.Pool, name string, price float64, symbol string) error {
	_, err := conn.Exec(context.Background(), `INSERT INTO stocks (name, price, symbol) VALUES ($1, $2, $3)`, name, price, symbol)
	if err != nil {
		return fmt.Errorf("unable to insert stock: %v", err)
	}
	return nil
}

// StockExists checks if a stock with the given name or symbol already exists.
func StockExists(conn *pgxpool.Pool, name, symbol string) (bool, error) {
	var exists bool
	log.Printf("Checking if stock exists: name=%s, symbol=%s", name, symbol)
	err := conn.QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM stocks WHERE name=$1 OR symbol=$2)`, name, symbol).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("unable to check if stock exists: %v", err)
	}
	log.Printf("Stock exists: %v", exists)
	return exists, nil
}
