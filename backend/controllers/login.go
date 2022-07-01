package controllers

import (
	"backend/database"
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var loginUser = models.User{}
	database.DB.Debug().Where("email = ?", data["email"]).First(&loginUser)

	// If email address is not found
	if loginUser.ID == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Email address not found")
	}

	// If email address is found
	// Compare password

	err = bcrypt.CompareHashAndPassword(loginUser.Password, []byte(data["password"]))
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Incorrect password, user is unauthorized")
	} else {
		return c.JSON(fiber.Map{"data": loginUser,
			"error":   err,
			"status":  fiber.StatusOK,
			"message": "User Validated"})
	}

	// Email found, password validated
	// Success message

}
