package routes

import (
	"todo/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(base *fiber.Group) error {
	auth := base.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)
	return nil
}
