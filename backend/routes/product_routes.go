package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(app *fiber.App) {
	app.Get("/api/products", controllers.GetAllProducts)
}

func ShowProduct(app *fiber.App) {
	app.Get("/api/product/:productID", controllers.ShowProduct)
}

func StoreProduct(app *fiber.App) {
	app.Post("/api/admin/product/create", controllers.StoreProduct)
}

func UpdateProduct(app *fiber.App) {
	app.Put("/api/admin/product/:productID", controllers.UpdateProduct)
}

func DeleteProduct(app *fiber.App) {
	app.Delete("/api/admin/product/:productID", controllers.DeleteProduct)
}