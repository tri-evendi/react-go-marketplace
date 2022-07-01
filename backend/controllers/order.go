package controllers

import (
	"backend/database"
	"backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type OrderDetails struct {
	OrderID     uint
	ProductID   uint
	ProductName string
	UserID      uint
	UserName    string
	Quantity    int64
	Amount      float64
	DateOrdered time.Time
	Status      string
	CreatedAt   time.Time
}

func GetAllOrders(c *fiber.Ctx) error {
	var orders = []models.Order{}
	database.DB.Preload("Product").Preload("User").Find(&orders)

	return c.JSON(fiber.Map{
		"data":   orders,
		"status": fiber.StatusOK,
	})
}

func ShowOrder(c *fiber.Ctx) error {
	orderID := c.Params("orderID")

	var order models.Order
	database.DB.Preload("Product").Preload("User").First(&order, orderID)

	return c.JSON(fiber.Map{
		"data":   order,
		"status": fiber.StatusOK,
	})
}

func validMultiIDs(productId_ uint, userId_ uint) (bool, string) {
	valid := true
	var _ error
	var exists bool
	var allerror []bool
	var errorstring string

	var names = [4]string{"ProductId"}

	_ = database.DB.Model(&models.Product{}).Select("count(*) > 0").Where("id = ?", productId_).Find(&exists).Error
	err1 := exists
	_ = database.DB.Model(&models.User{}).Select("count(*) > 0").Where("id = ?", userId_).Find(&exists).Error
	err2 := exists

	allerror = append(allerror, err1, err2)

	for i := 0; i < len(allerror); i++ {
		if !allerror[i] {
			valid = false
			errorstring = errorstring + names[i] + ", "
		}
	}
	return valid, errorstring
}

func StoreOrder(c *fiber.Ctx) error {
	data := new(models.Order)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Order failed.")
	}

	productId := data.ProductID
	userId := data.UserID
	quantity := data.Quantity
	amount := data.Amount
	dateOrdered := time.Now()
	status := "pending"
	createdAt := time.Now()

	productId_ := uint(productId)
	userId_ := uint(userId)

	// Check IDs are vaild
	validity, screwedIds := validMultiIDs(productId_, userId_)
	if !validity {
		return fiber.NewError(fiber.StatusNotFound, screwedIds+" doesn't exist")
	}

	order := models.Order{
		ProductID:   productId_,
		UserID:      userId_,
		Quantity:    quantity,
		Amount:      amount,
		DateOrdered: dateOrdered,
		Status:      status,
		CreatedAt:   createdAt,
	}

	result := database.DB.Create(&order)

	database.DB.Preload("Product").Preload("User").Where("product_id", order.ProductID).Find(&order)

	orderResult := OrderDetails{
		OrderID:     order.OrderID,
		ProductID:   order.ProductID,
		ProductName: order.Product.Name,
		UserID:      order.UserID,
		UserName:    order.User.FirstName,
		Quantity:    order.Quantity,
		Amount:      order.Amount,
		DateOrdered: order.DateOrdered,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt,
	}

	// Confirmation on order
	return c.JSON(fiber.Map{
		"data":          orderResult,
		"status":        fiber.StatusOK,
		"error":         result.Error,
		"rows_affected": result.RowsAffected})
}

func UpdateOrder(c *fiber.Ctx) error {
	data := new(models.Order)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Order failed.")
	}

	orderID := c.Params("orderID")

	productId := data.ProductID
	userId := data.UserID
	quantity := data.Quantity
	amount := data.Amount
	status := data.Status

	productId_ := uint(productId)
	userId_ := uint(userId)

	// Check IDs are vaild
	validity, screwedIds := validMultiIDs(productId_, userId_)
	if !validity {
		return fiber.NewError(fiber.StatusNotFound, screwedIds+" doesn't exist")
	}

	newOrder := models.Order{
		ProductID:   productId_,
		UserID:      userId_,
		Quantity:    quantity,
		Amount:      amount,
		Status:      status,
	}

	result := database.DB.Where("order_id", orderID).Updates(&newOrder)

	var afterUpdate models.Order
	database.DB.Preload("Product").Preload("User").First(&afterUpdate, orderID)

	orderResult := OrderDetails{
		OrderID:     afterUpdate.OrderID,
		ProductID:   afterUpdate.ProductID,
		ProductName: afterUpdate.Product.Name,
		UserID:      afterUpdate.UserID,
		UserName:    afterUpdate.User.FirstName,
		Quantity:    afterUpdate.Quantity,
		Amount:      afterUpdate.Amount,
		DateOrdered: afterUpdate.DateOrdered,
		Status:      afterUpdate.Status,
		CreatedAt:   afterUpdate.CreatedAt,
	}

	// Confirmation on order
	return c.JSON(fiber.Map{
		"data":             orderResult,
		"status":           fiber.StatusOK,
		"error":      		result.Error,
		"rows_affected":    result.RowsAffected,
	})
}

func DeleteOrder(c *fiber.Ctx) error {
	orderID := c.Params("orderID")

	var order models.Order
	database.DB.First(&order, orderID)

	// Delete an existing record
	database.DB.Delete(&order)

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})
}


func MultipleStoreOrder(c *fiber.Ctx) error {
	data := new([]models.Order)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Order failed.")
	}

	var orders []models.Order
	var orderResult []OrderDetails
	var orderDetails OrderDetails
	var order models.Order
	var dateOrdered time.Time
	var createdAt time.Time

	for i := 0; i < len(*data); i++ {
		productId := (*data)[i].ProductID
		userId := (*data)[i].UserID
		quantity := (*data)[i].Quantity
		amount := (*data)[i].Amount
		dateOrdered = time.Now()
		status := "pending"
		createdAt = time.Now()

		productId_ := uint(productId)
		userId_ := uint(userId)

		// Check IDs are vaild
		validity, screwedIds := validMultiIDs(productId_, userId_)
		if !validity {
			return fiber.NewError(fiber.StatusNotFound, screwedIds+" doesn't exist")
		}

		order = models.Order{
			ProductID:   productId_,
			UserID:      userId_,
			Quantity:    quantity,
			Amount:      amount,
			DateOrdered: dateOrdered,
			Status:      status,
			CreatedAt:   createdAt,
		}

		orders = append(orders, order)
	}

	result := database.DB.Create(&orders)

	for i := 0; i < len(orders); i++ {
		database.DB.Preload("Product").Preload("User").Where("product_id", orders[i].ProductID).Find(&order)

		orderDetails = OrderDetails{
			OrderID:     order.OrderID,
			ProductID:   order.ProductID,
			ProductName: order.Product.Name,
			UserID:      order.UserID,
			UserName:    order.User.FirstName,
			Quantity:    order.Quantity,
			Amount:      order.Amount,
			DateOrdered: order.DateOrdered,
			Status:      order.Status,
			CreatedAt:   order.CreatedAt,
		}

		orderResult = append(orderResult, orderDetails)
	}

	// Confirmation on order
	return c.JSON(fiber.Map{
		"data":          orderResult,
		"status":        fiber.StatusOK,
		"error":         result.Error,
		"rows_affected": result.RowsAffected})
}
