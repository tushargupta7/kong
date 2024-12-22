package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tushargupta7/kong/errors"
	"github.com/tushargupta7/kong/handlers"
	"github.com/tushargupta7/kong/repositories"
	"github.com/tushargupta7/kong/routes"
	"log"
	"os"
)
import database "github.com/tushargupta7/kong/database"

func main() {
	// Initialize database connection
	if err := database.InitDB(); err != nil {
		// Log the error and stop the application if there's a database connection issue
		log.Fatalf("Error initializing database: %v", err)
		os.Exit(1) // Exit with a non-zero code to indicate failure
	}

	// Initialize the service repository
	serviceRepo := repositories.NewServiceRepository(database.DB)

	// Initialize the service handler with the repository
	handlers.InitServiceHandler(serviceRepo)

	// Initialize Fiber
	app := fiber.New()
	app.Use(errorHandler)

	// Register routes
	routes.RegisterServiceRoutes(app)
	routes.RegisterVersionRoutes(app)
	routes.RegisterLoginRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}

func errorHandler(c *fiber.Ctx) error {
	// Check if the error is of type AppError
	if err := c.Next(); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			// If it's an AppError, return the response based on AppError properties
			return c.Status(appErr.StatusCode).JSON(fiber.Map{
				"error":   appErr.Message,
				"code":    appErr.Err,
				"details": appErr.Context, // Include context data if necessary
			})
		}

		// If it's not an AppError, return a generic 500 internal server error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return nil
}
