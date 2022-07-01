package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetAllOrders(app *fiber.App) {
	app.Get("/api/admin/orders", controllers.GetAllOrders)
}

func ShowOrder(app *fiber.App) {
	app.Get("/api/admin/order/:orderID", controllers.ShowOrder)
}

func StoreOrder(app *fiber.App) {
	app.Post("/api/order/create", controllers.StoreOrder)
}

func MultipleStoreOrder(app *fiber.App) {
	app.Post("/api/order/bulk-create", controllers.MultipleStoreOrder)
}

func UpdateOrder(app *fiber.App) {
	app.Put("/api/admin/order/:orderID", controllers.UpdateOrder)
}

func DeleteOrder(app *fiber.App) {
	app.Delete("/api/admin/order/:orderID", controllers.DeleteOrder)
}