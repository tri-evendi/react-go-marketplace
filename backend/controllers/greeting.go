package controllers

import "github.com/gofiber/fiber/v2"

func Greeting(c *fiber.Ctx) error {
	return c.SendString("Hello, Worlds!")
}
