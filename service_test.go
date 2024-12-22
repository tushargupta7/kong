package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tushargupta7/kong/database"
	"github.com/tushargupta7/kong/dtos"
	"github.com/tushargupta7/kong/middleware"
	"github.com/tushargupta7/kong/repositories"
	"github.com/tushargupta7/kong/routes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// Helper function to initialize Fiber app for testing
func setupApp() *fiber.App {
	app := fiber.New()

	// Register the service routes
	routes.RegisterServiceRoutes(app)
	routes.RegisterLoginRoutes(app)
	// Mock JWT middleware (for testing)
	app.Use(middleware.JWTMiddleware("admin"))
	return app
}

// Helper function to login and fetch a token
func loginAndGetToken(t *testing.T, app *fiber.App) string {
	// Test login credentials
	loginRequest := dtos.LoginRequest{
		Username: "admin",
		Password: "Admin@123",
	}

	// Convert request data to JSON
	jsonData, err := json.Marshal(loginRequest)
	assert.Nil(t, err)

	// Send POST request to login
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(jsonData))
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, 200)

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	// Extract the token
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	token, ok := response["token"].(string)
	assert.True(t, ok)

	return token
}

// Test creating a new service with token
func TestCreateService(t *testing.T) {
	app := setupApp()

	// Login and get token
	token := loginAndGetToken(t, app)

	// Test data
	serviceRequest := dtos.CreateServiceRequest{
		Name:        "Test Service",
		Description: "This is a test service",
	}

	// Convert request data to JSON
	jsonData, err := json.Marshal(serviceRequest)
	assert.Nil(t, err)

	// Send POST request to create service
	req := httptest.NewRequest(http.MethodPost, "/service", bytes.NewReader(jsonData))
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)

	// Assert the status and response
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, 201) // HTTP Status Code Created

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	// Optionally check if the response contains the correct service data
	var serviceResponse map[string]interface{}
	err = json.Unmarshal(body, &serviceResponse)
	assert.Nil(t, err)
	assert.Equal(t, serviceResponse["name"], serviceRequest.Name)
	assert.Equal(t, serviceResponse["description"], serviceRequest.Description)
}

// Test viewing a created service
func TestGetService(t *testing.T) {
	app := setupApp()

	// Create a service (you can also do this within the test if needed)
	service := dtos.CreateServiceRequest{
		Name:        "Test Service",
		Description: "This is a test service",
	}
	createdService, _ := repositories.InsertService(database.DB, service.Name, service.Description)

	// Convert ID to string for URL
	serviceID := strconv.Itoa(int(createdService.ID))

	// Send GET request to fetch the service by ID
	req := httptest.NewRequest(http.MethodGet, "/service/"+serviceID, nil)
	resp, err := app.Test(req)

	// Assert the status and response
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, 200) // HTTP Status Code OK

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	// Optionally check if the response contains the correct service data
	var serviceResponse map[string]interface{}
	err = json.Unmarshal(body, &serviceResponse)
	assert.Nil(t, err)
	assert.Equal(t, serviceResponse["name"], service.Name)
	assert.Equal(t, serviceResponse["description"], service.Description)
}

// Test deleting a service
func TestDeleteService(t *testing.T) {
	app := setupApp()

	// Create a service (you can also do this within the test if needed)
	service := dtos.CreateServiceRequest{
		Name:        "Test Service",
		Description: "This is a test service",
	}
	createdService, _ := repositories.InsertService(database.DB, service.Name, service.Description)

	// Convert ID to string for URL
	serviceID := strconv.Itoa(int(createdService.ID))

	// Send DELETE request to remove the service by ID
	req := httptest.NewRequest(http.MethodDelete, "/service/"+serviceID, nil)
	resp, err := app.Test(req)

	// Assert the status and response
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, 200) // HTTP Status Code OK

	// Optionally, you could verify the service was actually deleted by attempting to fetch it again
	_, err = repositories.GetServiceByID(database.DB, strconv.Itoa(int(createdService.ID)))
	assert.Equal(t, err, sql.ErrNoRows) // Expecting error because the service should not exist
}
