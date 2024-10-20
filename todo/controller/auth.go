package controller

import (
	"github.com/gofiber/fiber/v2"

	"todo/db"
	"todo/util"
)

type RegisterData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	c.SendString("I'm a POST request!")
	return nil
}

func Register(c *fiber.Ctx) error {
	var formated_data RegisterData

	if err := c.BodyParser(&formated_data); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	email := formated_data.Email
	password := formated_data.Password

	session, err := db.DBConnection()
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Failed to connect to database",
		})
	}

	hashed_password, err := util.HashPassword(password)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	session.Create(db.User{Email: email, Password: hashed_password})
	return c.JSON(fiber.Map{
		"message": "Registered Successfully",
	})

}
