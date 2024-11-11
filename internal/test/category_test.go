package test

import (
	"api-payment/internal/entity"
	"api-payment/internal/router"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	//"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/api-payment_db")
}

func TestCategory(t *testing.T) {

	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	defer db.Close()

	//router := initializeRoute(db)

	t.Run("get Categories", func(t *testing.T) {

		router := router.InitializeRoute(db)

		req := httptest.NewRequest(http.MethodGet, "/category", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)

		var categories []map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&categories)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(categories), 0)
	})

	t.Run("create category", func(t *testing.T) {

		router := router.InitializeRoute(db)

		category := entity.NewCategory("Books")

		fmt.Println("category:", category)
		// Convert category to JSON
		categoryJSON, err := json.Marshal(category)
		if err != nil {
			t.Fatalf("Failed to marshal category: %v", err)
		}

		fmt.Println("Json:", string(categoryJSON))

		req := httptest.NewRequest(http.MethodGet, "/category", strings.NewReader(string(categoryJSON)))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusCreated, resp.Code)

		var cat entity.Category
		err = json.NewDecoder(resp.Body).Decode(&cat)
		assert.NoError(t, err)
		assert.Equal(t, "Books", cat.Name)

	})

}
