package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// ConnectDB ansluter till PostgreSQL och returnerar en anslutning.
func ConnectDB() (*pgx.Conn, error) {
	connStr := "postgres://postgres:1234@localhost:5433/stock_analysis_db?sslmode=disable"
	return pgx.Connect(context.Background(), connStr)
}
