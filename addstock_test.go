package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-stock-analysis/database"

	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/assert"
)

func Test_When_Adding_A_New_Stock(t *testing.T) {
	// Arrange
	mockDB := &database.MockDB{
		ExecFunc: func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, nil
		},
	}
	router := setupRouter(mockDB)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/addstock", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Stock added")
}
