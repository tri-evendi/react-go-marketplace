package controllers

import (
	"backend/config"
	"backend/database"
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), config.PasswordCost)

	user := models.User{
		FirstName: data["firstName"],
		LastName:  data["lastName"],
		Email:     data["email"],
		Password:  password,
	}

	result := database.DB.Create(&user)

	return c.JSON(fiber.Map{"data": user, "error": result.Error, "rows_affected": result.RowsAffected})
}
