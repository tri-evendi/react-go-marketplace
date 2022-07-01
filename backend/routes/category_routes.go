package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(app *fiber.App) {
	app.Get("/api/categories", controllers.GetAllCategories)
}

func GetProductByCategory(app *fiber.App) {
	app.Get("/api/category/:categoryID/products", controllers.GetProductByCategory)
}

func ShowCategory(app *fiber.App) {
	app.Get("/api/admin/category/:categoryID", controllers.ShowCategory)
}

func StoreCategory(app *fiber.App) {
	app.Post("/api/admin/category/create", controllers.StoreCategory)
}

func UpdateCategory(app *fiber.App) {
	app.Put("/api/admin/category/:categoryID", controllers.UpdateCategory)
}

func DeleteCategory(app *fiber.App) {
	app.Delete("/api/admin/category/:categoryID", controllers.DeleteCategory)
}