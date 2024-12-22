package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tushargupta7/kong/handlers"
	"github.com/tushargupta7/kong/middleware"
)

// RegisterServiceRoutes registers routes for service-related operations
func RegisterServiceRoutes(app *fiber.App) {
	serviceGroup := app.Group("/service")

	serviceGroup.Post("/", middleware.JWTMiddleware("admin"), handlers.CreateService)      // POST /services
	serviceGroup.Put("/:id", middleware.JWTMiddleware("admin"), handlers.UpdateService)    // PUT /services/:id
	serviceGroup.Delete("/:id", middleware.JWTMiddleware("admin"), handlers.DeleteService) // DELETE /services/:id

	serviceGroup.Get("/", middleware.JWTMiddleware("user"), handlers.GetServices)   // GET /services
	serviceGroup.Get("/:id", middleware.JWTMiddleware("user"), handlers.GetService) // GET /services/:id
}
