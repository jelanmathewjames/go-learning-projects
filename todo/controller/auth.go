package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	c.SendString("I'm a POST request!")
	return nil
}

func Register(c *fiber.Ctx) error {
	var formated_data map[string]interface{}

	if err := c.BodyParser(&formated_data); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Registered Successfully",
	})

}
