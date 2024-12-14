package routes

import (
	"todo/controller"
	"todo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(base *fiber.Group) {
	todo := base.Group("/todo")
	todo.Use(middlewares.AuthMiddleware)
	todo.Get("/", controller.ListTodos)
	todo.Post("/", controller.CreateTodo)
	todo.Patch("/:id", controller.UpdateTodo)
	todo.Delete("/:id", controller.DeleteTodo)
}
