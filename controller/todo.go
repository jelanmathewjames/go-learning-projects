package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"todo/db"
)

type TodoCreateData struct {
	Title  string    `json:"title"`
	Expiry time.Time `json:"expiry"`
}

func ListTodos(c *fiber.Ctx) error {
	user_id := c.Locals("user_data").(map[string]interface{})["user_id"].(string)
	todos := []db.Todo{}
	session, err := db.DBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	if result := session.Where("user_id = ?", user_id).Find(&todos); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No Todos Found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "List of Todos",
		"data":    todos,
	})
}

func CreateTodo(c *fiber.Ctx) error {
	var todo_data TodoCreateData
	user_id := c.Locals("user_data").(map[string]interface{})["user_id"].(string)
	if err := c.BodyParser(&todo_data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	session, err := db.DBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	session.Create(&db.Todo{
		Title:  todo_data.Title,
		Expiry: todo_data.Expiry,
		UserID: user_id,
	})
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo Created",
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	var todo_data map[string]interface{}
	user_id := c.Locals("user_data").(map[string]interface{})["user_id"].(string)
	todo_id := c.Params("id")
	if err := c.BodyParser(&todo_data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	session, err := db.DBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	var todo db.Todo

	if err := session.Where("id = ? AND user_id = ?", todo_id, user_id).First(&todo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	if err := session.Model(&todo).Updates(todo_data).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo Updated",
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	var todo db.Todo
	user_id := c.Locals("user_data").(map[string]interface{})["user_id"].(string)
	todo_id := c.Params("id")
	session, err := db.DBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	if err := session.Where("id = ? AND user_id = ?", todo_id, user_id).Delete(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting todo",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo Deleted",
	})
}
