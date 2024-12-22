package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tushargupta7/kong/handlers"
)

// SetupRoutes initializes the routes for the application
func RegisterLoginRoutes(app *fiber.App) {
	// Public route for login
	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.Login(c)
	})
}
