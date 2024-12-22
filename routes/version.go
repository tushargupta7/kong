package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tushargupta7/kong/handlers"
	"github.com/tushargupta7/kong/middleware"
)

func RegisterVersionRoutes(app *fiber.App) {
	serviceGroup := app.Group("/service/:id/version")

	serviceGroup.Post("/", middleware.JWTMiddleware("admin"), handlers.CreateVersion)
	serviceGroup.Delete("/:versionId", middleware.JWTMiddleware("admin"), handlers.DeleteVersion)
	serviceGroup.Put("/:versionId", middleware.JWTMiddleware("admin"), handlers.UpdateVersion)
	serviceGroup.Get("/", middleware.JWTMiddleware("user"), handlers.GetServiceVersions)
}
