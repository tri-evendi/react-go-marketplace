package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetAllCodes(app *fiber.App) {
	app.Get("/api/admin/codes", controllers.GetAllCodes)
}

func ShowCode(app *fiber.App) {
	app.Get("/api/admin/code/:codeID", controllers.ShowCode)
}

func StoreCode(app *fiber.App) {
	app.Post("/api/admin/code/create", controllers.StoreCode)
}

func UpdateCode(app *fiber.App) {
	app.Put("/api/admin/code/:codeID", controllers.UpdateCode)
}

func DeleteCode(app *fiber.App) {
	app.Delete("/api/admin/code/:codeID", controllers.DeleteCode)
}