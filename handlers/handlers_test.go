package handlers

import (
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tushargupta7/kong/database"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetServices(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	// Override the global DB variable
	database.DB = db

	// Mock query
	mock.ExpectQuery(`SELECT id, name, description, created_at, updated_at FROM services WHERE name ILIKE \$1 ORDER BY name asc LIMIT \$2 OFFSET \$3`).
		WithArgs("%serv%", 10, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
			AddRow(1, "Service 1", "Description 1", time.Now(), time.Now()).
			AddRow(2, "Service 2", "Description 2", time.Now(), time.Now()))

	// Create Fiber app
	app := fiber.New()

	// Register the handler
	app.Get("/service", GetServices)

	// Perform the request
	req := httptest.NewRequest("GET", "/service?search=serv&limit=10&page=1", nil)
	resp, err := app.Test(req, -1) // No timeout
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	// Assert the status code
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Parse response body
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)

	// Validate response structure
	assert.Equal(t, 1.0, result["page"]) // Adjust if needed
	assert.Equal(t, 10.0, result["limit"])
	services := result["results"].([]interface{})
	assert.Len(t, services, 2)
}

func TestGetService(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	// Override the global DB variable for testing
	database.DB = db

	// Define the mock behavior for a service
	mock.ExpectQuery(`SELECT id, name, description, created_at, updated_at FROM services WHERE id = \$1`).
		WithArgs(int64(1)). // Use int64 for the ID
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
			AddRow(1, "Service 1", "Description 1", time.Now(), time.Now())) // Include all columns

	// Create a new Fiber app
	app := fiber.New()

	// Register the handler
	app.Get("/service/:id", GetService)

	// Perform the request
	req := httptest.NewRequest("GET", "/service/1", nil)
	resp, err := app.Test(req, 100000)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	// Assert the status code is 200 OK
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Check the response body
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)

	// Assert the returned service details
	assert.Equal(t, 1.0, result["id"]) // Fiber parses numbers as float64
	assert.Equal(t, "Service 1", result["name"])
	assert.Equal(t, "Description 1", result["description"])
}

//Test for not found etc..
