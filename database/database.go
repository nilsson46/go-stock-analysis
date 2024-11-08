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
	// Create table if not exists
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

	// Update table schema if needed
	_, err = conn.Exec(context.Background(), `ALTER TABLE stocks
        ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        ADD COLUMN IF NOT EXISTS direktavkastning DECIMAL,
        ADD COLUMN IF NOT EXISTS ep DECIMAL,
        ADD COLUMN IF NOT EXISTS egetkapital DECIMAL`)
	if err != nil {
		log.Fatalf("Unable to update table schema: %v\n", err)
	}

	// Insert a stock entry
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
func AddStock(conn *pgxpool.Pool, name string, kurs float64, symbol string) error {
	_, err := conn.Exec(context.Background(), `INSERT INTO stocks (name, kurs, symbol) VALUES ($1, $2, $3)`, name, kurs, symbol)
	if err != nil {
		return fmt.Errorf("unable to insert stock: %v", err)
	}
	return nil
}
