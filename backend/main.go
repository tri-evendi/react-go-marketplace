package main

import (
	"backend/database"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	//DataBase Connection Setup
	database.Connect()

	// App object creation
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Static("/", "./img")

	// Add "/" route
	routes.Greeting(app)
	// Add "/api/register" route
	routes.Register(app)
	// Add "/api/login" route
	routes.Login(app)

	// Add "/api/products" route
	routes.GetAllProducts(app)
	routes.ShowProduct(app)
	routes.StoreProduct(app)
	routes.UpdateProduct(app)
	routes.DeleteProduct(app)

	// Get all products by category
	routes.GetProductByCategory(app)

	// Add "/api/categories" route
	routes.GetAllCategories(app)
	routes.ShowCategory(app)
	routes.StoreCategory(app)
	routes.UpdateCategory(app)
	routes.DeleteCategory(app)

	// Add "/api/codes" route
	routes.GetAllCodes(app)
	routes.ShowCode(app)
	routes.StoreCode(app)
	routes.UpdateCode(app)
	routes.DeleteCode(app)

	// Add "/api/orders" route
	routes.GetAllOrders(app)
	routes.ShowOrder(app)
	routes.StoreOrder(app)
	routes.MultipleStoreOrder(app)
	routes.UpdateOrder(app)
	routes.DeleteOrder(app)

	// Add "/api/payments" route
	routes.GetAllPayments(app)
	routes.ShowPayment(app)
	routes.StorePayment(app)
	routes.UpdatePayment(app)
	routes.DeletePayment(app)

	app.Listen(":8080")
}