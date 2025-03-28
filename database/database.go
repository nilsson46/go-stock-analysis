// database/database.go
package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Close()
}

type SQLDB struct {
	*pgxpool.Pool
}

func (db *SQLDB) Close() {
	db.Pool.Close()
}

// ConnectDB ansluter till PostgreSQL-databasen
func ConnectDB() (*SQLDB, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatalf("DATABASE_URL is not set")
	}
	conn, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}
	return &SQLDB{conn}, nil
}

// InitializeDB initialiserar databasen
func InitializeDB(conn DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := conn.Exec(ctx, `
        CREATE TABLE IF NOT EXISTS stocks (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            price DECIMAL(10, 2) NOT NULL,
            symbol VARCHAR(10) NOT NULL
        )
    `)
	if err != nil {
		log.Printf("Unable to create table: %v\n", err)
	}
}

// StockExists kontrollerar om en aktie med samma namn eller symbol redan finns i databasen
func StockExists(conn DB, name, symbol string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM stocks WHERE name=$1 OR symbol=$2)`
	err := conn.QueryRow(context.Background(), query, name, symbol).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if stock exists: %v", err)
	}
	return exists, err
}

// AddStock lägger till en ny aktie i databasen
func AddStock(conn DB, name string, price float64, symbol string) error {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO stocks (name, price, symbol) VALUES ($1, $2, $3)`,
		name, price, symbol,
	)
	if err != nil {
		log.Printf("Error adding stock: %v", err)
		return err
	}
	return nil
}

// GetStock hämtar en specifik aktie baserat på dess namn eller symbol
func GetStock(conn DB, name, symbol string) (map[string]interface{}, error) {
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
func GetStocksFromDB(conn DB) ([]map[string]interface{}, error) {
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
func DeleteStockBySymbol(db DB, symbol string) error {
	query := "DELETE FROM stocks WHERE symbol = $1"
	_, err := db.Exec(context.Background(), query, symbol)
	return err
}

// UpdateStockPrice uppdaterar priset på en aktie baserat på dess symbol
func UpdateStockPrice(db DB, symbol string, newPrice float64) error {
	query := "UPDATE stocks SET price = $1 WHERE symbol = $2"
	_, err := db.Exec(context.Background(), query, newPrice, symbol)
	if err != nil {
		log.Printf("Error updating stock price: %v", err)
		return err
	}
	return nil
}
