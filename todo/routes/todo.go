package routes

import (
	"todo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(base *fiber.Group) {
	todo := base.Group("/todo")
	todo.Use(middlewares.AuthMiddleware)
	todo.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	// todo.POST("/", create)
	// todo.GET("/", list)
	// todo.GET("/:id", get)
	// todo.PUT("/:id", update)
	// todo.DELETE("/:id", delete)
}
