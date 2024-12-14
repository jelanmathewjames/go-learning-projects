package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	base_api := app.Group("/api/v1")
	AuthRoutes(base_api.(*fiber.Group))
	TodoRoutes(base_api.(*fiber.Group))
}
