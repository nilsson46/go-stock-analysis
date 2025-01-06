package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-stock-analysis/database"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func Test_When_Adding_A_New_Stock(t *testing.T) {
	// Arrange
	mockDB := &database.MockDB{
		QueryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return &database.MockRow{
				ScanFunc: func(dest ...interface{}) error {
					*dest[0].(*bool) = false
					return nil
				},
			}
		},
		ExecFunc: func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, nil
		},
	}
	router := setupRouter(mockDB)

	// Skapa en JSON-payload för POST-begäran
	stock := map[string]interface{}{
		"name":   "Test Stock",
		"price":  100.0,
		"symbol": "TST",
	}
	jsonValue, _ := json.Marshal(stock)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/addstock", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Stock added")
}
