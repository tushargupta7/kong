package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tushargupta7/kong/errors"
	"github.com/tushargupta7/kong/handlers"
	"github.com/tushargupta7/kong/migrations"
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

	// Run migrations
	migrations.Migrate(database.DB)

	// Initialize the service repository , similar to be done for other entities
	serviceRepo := repositories.NewServiceRepository(database.DB)

	// Initialize the service handler with the repository
	handlers.InitServiceHandler(serviceRepo)

	// Initialize Fiber
	app := fiber.New()
	app.Use(errors.ErrorHandler)

	// Register routes
	routes.RegisterServiceRoutes(app)
	routes.RegisterVersionRoutes(app)
	routes.RegisterLoginRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
