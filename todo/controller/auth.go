package controller

import (
	"github.com/gofiber/fiber/v2"

	"todo/db"
	"todo/util"
)

type UserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var formated_data UserData
	var user db.User
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
	if result := session.Where("email = ?", email).First(&user); result.Error == nil {
		if err := util.ComparePassword(password, user.Password); err != nil {
			return c.JSON(fiber.Map{
				"message": "Invalid Password",
			})
		}
		if token, err := util.GenerateToken(user.ID); err == nil {
			return c.JSON(fiber.Map{
				"message": "Login Successful",
				"token":   token,
			})
		}
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})

	}
	return c.JSON(fiber.Map{
		"message": "User not found",
	})
}

func Register(c *fiber.Ctx) error {
	var formated_data UserData

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

	session.Create(&db.User{Email: email, Password: hashed_password})
	return c.JSON(fiber.Map{
		"message": "Registered Successfully",
	})

}
