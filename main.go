package kong

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tushargupta7/kong/db"
	"github.com/tushargupta7/kong/handlers"
	"log"
)

func main() {
	// Initialize database connection
	db.InitDB()

	// Initialize Fiber
	app := fiber.New()

	// Routes
	app.Get("/services", handlers.GetServices)
	app.Get("/services/:id", handlers.GetService)
	app.Get("/services/:id/versions", handlers.GetServiceVersions)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
