package controllers

import (
	"backend/database"
	"backend/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type PaymentDetails struct {
	PaymentID   uint
	OrderID     uint
	UserID      uint
	UserName    string
	Type        string
	Amount      float64
	DatePayment time.Time
	Status      string
	CreatedAt   time.Time
}

func GetAllPayments(c *fiber.Ctx) error {
	var payments = []models.Payment{}
	database.DB.Preload("Order").Preload("User").Find(&payments)

	return c.JSON(fiber.Map{
		"data":   payments,
		"status": fiber.StatusOK,
	})
}

func ShowPayment(c *fiber.Ctx) error {
	paymentID := c.Params("paymentID")

	var payment models.Payment
	database.DB.Preload("Order").Preload("User").First(&payment, paymentID)

	return c.JSON(fiber.Map{
		"data":   payment,
		"status": fiber.StatusOK,
	})
}

func validPaymentMultiIDs(orderId_ uint, userId_ uint) (bool, string) {
	valid := true
	var _ error
	var exists bool
	var allerror []bool
	var errorstring string

	var names = [4]string{"OrderId"}

	_ = database.DB.Model(&models.Order{}).Select("count(*) > 0").Where("order_id = ?", orderId_).Find(&exists).Error
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

func StorePayment(c *fiber.Ctx) error {
	data := new(models.Payment)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Payment failed.")
	}

	orderID := c.Params("orderID")
	

	orderId, err1 := strconv.Atoi(orderID)
	userId := data.UserID
	types := data.Type
	amount := data.Amount
	datePaymented := time.Now()
	status := data.Status
	createdAt := time.Now()

	if err1 != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, "Data Validation failed. Payment failed.")
	}

	orderId_ := uint(orderId)
	userId_ := uint(userId)

	// Check IDs are vaild
	validity, screwedIds := validPaymentMultiIDs(orderId_, userId_)
	if !validity {
		return fiber.NewError(fiber.StatusNotFound, screwedIds+" doesn't exist")
	}

	payment := models.Payment{
		OrderID:     orderId_,
		UserID:      userId_,
		Type:        types,
		Amount:      amount,
		DatePayment: datePaymented,
		Status:      status,
		CreatedAt:   createdAt,
	}

	result := database.DB.Create(&payment)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Payment failed.")
	} else {
		var orderDetail models.Order
		database.DB.Preload("Product").Preload("User").First(&orderDetail, orderID)
		// Return the created record

		if payment.Status == "success" {
			var codes []models.Code
			database.DB.Limit(int(orderDetail.Quantity)).Where("product_id = ?", orderDetail.ProductID).Find(&codes)
			
			var datas []uint
			for _, code := range codes {
				datas = append(datas, uint(code.ID))
			}

			for i := 0; i < len(datas); i++ {
				coderable := models.Coderable{
					OrderID:     orderId_,
					CodeID:      datas[i],
				}
				database.DB.Create(&coderable)
			}
	
			database.DB.Preload("Order").Where("order_id = ?", payment.OrderID).Find(&payment)
		
			paymentResult := PaymentDetails{
				PaymentID:   payment.PaymentID,
				OrderID:     payment.OrderID,
				UserID:      payment.UserID,
				UserName:    payment.User.FirstName,
				Type:        payment.Type,
				Amount:      payment.Amount,
				DatePayment: payment.DatePayment,
				Status:      payment.Status,
				CreatedAt:   payment.CreatedAt,
			}
		
			return c.JSON(fiber.Map{
				"data":          paymentResult,
				"status":        fiber.StatusOK,
				"error":         result.Error,
				"rows_affected": result.RowsAffected})
		}
		return c.JSON(fiber.Map{
			"error":         result.Error,
			"rows_affected": result.RowsAffected})
	}

}

func UpdatePayment(c *fiber.Ctx) error {
	data := new(models.Payment)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Payment failed.")
	}

	orderID := c.Params("orderID")
	paymentID := c.Params("paymentID")

	orderId, err1 := strconv.Atoi(orderID)
	userId := data.UserID
	types := data.Type
	amount := data.Amount
	status := data.Status

	if err1 != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, "Data Validation failed. Payment failed.")
	}

	orderId_ := uint(orderId)
	userId_ := uint(userId)

	newPayment := models.Payment{
		UserID:  userId_,
		Type:    types,
		Amount:  amount,
		Status:  status,
	}

	result := database.DB.Where("payment_id", paymentID).Updates(&newPayment)

	var orderDetail models.Order
	database.DB.Preload("Product").Preload("User").First(&orderDetail, orderId)

	if newPayment.Status == "success" {
		var codes []models.Code
		database.DB.Limit(int(orderDetail.Quantity)).Where("product_id = ?", orderDetail.ProductID).Find(&codes)
		
		var datas []uint
		for _, code := range codes {
			datas = append(datas, uint(code.ID))
		}

		for i := 0; i < len(datas); i++ {
			coderable := models.Coderable{
				OrderID:     orderId_,
				CodeID:      datas[i],
			}
			database.DB.Create(&coderable)
		}

		var afterUpdate models.Payment
		database.DB.First(&afterUpdate, paymentID)

		paymentResult := PaymentDetails{
			PaymentID:   afterUpdate.PaymentID,
			OrderID:     afterUpdate.OrderID,
			UserID:      afterUpdate.UserID,
			UserName:    afterUpdate.User.FirstName,
			Type:        afterUpdate.Type,
			Amount:      afterUpdate.Amount,
			DatePayment: afterUpdate.DatePayment,
			Status:      afterUpdate.Status,
			CreatedAt:   afterUpdate.CreatedAt,
		}

		return c.JSON(fiber.Map{
			"data":          paymentResult,
			"status":        fiber.StatusOK,
			"error":         result.Error,
			"rows_affected": result.RowsAffected,
		})
	}
	return c.JSON(fiber.Map{
		"error":         result.Error,
		"rows_affected": result.RowsAffected,
	})
}

func DeletePayment(c *fiber.Ctx) error {
	paymentID := c.Params("paymentID")

	var payment models.Payment
	database.DB.First(&payment, paymentID)

	// Delete an existing record
	database.DB.Delete(&payment)

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})

}
