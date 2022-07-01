package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetAllPayments(app *fiber.App) {
	app.Get("/api/admin/payments", controllers.GetAllPayments)
}

func ShowPayment(app *fiber.App) {
	app.Get("/api/admin/payment/:paymentID", controllers.ShowPayment)
}

func StorePayment(app *fiber.App) {
	app.Post("/api/order/:orderID/payment/create", controllers.StorePayment)
}

func UpdatePayment(app *fiber.App) {
	app.Put("/api/admin/:orderID/payment/:paymentID", controllers.UpdatePayment)
}

func DeletePayment(app *fiber.App) {
	app.Delete("/api/admin/payment/:paymentID", controllers.DeletePayment)
}