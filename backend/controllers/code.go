package controllers

import (
	"backend/database"
	"backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CodeDetails struct {
	ID          uint
	Code        string
	ProductId  	uint
	ProductName string
	DateExpired time.Time
	IsAvailable bool
}


func GetAllCodes(c *fiber.Ctx) error {
	var codes = []models.Code{}
	database.DB.Preload("Product").Find(&codes)

	return c.JSON(fiber.Map{
		"data":   codes,
		"status": fiber.StatusOK,
	})
}

func ShowCode(c *fiber.Ctx) error {
	codeID := c.Params("codeID")
	
	var code models.Code
	database.DB.Preload("Product").First(&code, codeID)

	return c.JSON(fiber.Map{
		"data":   code,
		"status": fiber.StatusOK,
	})
}

func validProductIDs( productId_ uint) (bool, string) {
	valid := true
	var _ error
	var exists bool
	var allerror []bool
	var errorstring string

	var names = [4]string{"ProductId"}

	_ = database.DB.Model(&models.Product{}).Select("count(*) > 0").Where("id = ?", productId_).Find(&exists).Error
	err1 := exists

	allerror = append(allerror, err1)

	for i := 0; i < len(allerror); i++ {
		if !allerror[i] {
			valid = false
			errorstring = errorstring + names[i] + ", "
		}
	}
	return valid, errorstring
}

func StoreCode(c *fiber.Ctx) error {
	data := new(models.Code)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Code failed.")
	}

	productId := data.ProductID
	nameCode 	:= data.Code
	dateExpiredCode 	:= time.Date(2022, 8, 1, 12, 13, 14, 0, time.UTC)
	isAvailable := data.IsAvailable

	productId_ := uint(productId)

	// Check IDs are vaild
	validity, screwedIds := validProductIDs(productId_)
	if !validity {
		return fiber.NewError(fiber.StatusNotFound, screwedIds+" doesn't exist")
	}

	code := models.Code{
		ProductID:     	productId_,
		Code:     		nameCode,
		DateExpired: 	dateExpiredCode,
		IsAvailable:    isAvailable ,
	}

	result := database.DB.Create(&code)

	database.DB.Preload("Product").Where("product_id = ?", code.ProductID).Find(&code)

	codeResult := CodeDetails{
		ID:      			code.ID,
		Code:   			code.Code,
		DateExpired: 		code.DateExpired,
		ProductId:         	code.ProductID,
		ProductName:       	code.Product.Name,
		IsAvailable:        code.IsAvailable,
	}

	// Confirmation on code
	return c.JSON(fiber.Map{
		"data":             codeResult,
		"status":           fiber.StatusOK,
		"error":      		result.Error,
		"rows_affected":    result.RowsAffected})
}

func UpdateCode(c *fiber.Ctx) error {
	data := new(models.Code)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Code failed.")
	}

	codeID := c.Params("codeID")

	productId := data.ProductID
	nameCode 	:= data.Code
	dateExpiredCode 	:=   time.Date(2022, 8, 1, 12, 13, 14, 0, time.UTC)
	isAvailable := data.IsAvailable

	productId_ := uint(productId)

	// Check IDs are vaild
	validity, screwedIds := validProductIDs(productId_)
	if !validity {
		return fiber.NewError(fiber.StatusNotFound, screwedIds+" doesn't exist")
	}

	newCode := models.Code{
		ProductID:     	productId_,
		Code:     		nameCode,
		DateExpired: 	dateExpiredCode,
		IsAvailable:    isAvailable ,
	}

	result := database.DB.Where("ID", codeID).Updates(&newCode)

	var afterUpdate models.Code
	database.DB.First(&afterUpdate, codeID)

	codeResult := CodeDetails{
		ID:      			afterUpdate.ID,
		Code:   			afterUpdate.Code,
		DateExpired: 		afterUpdate.DateExpired,
		ProductId:         	afterUpdate.ProductID,
		ProductName:       	afterUpdate.Product.Name,
		IsAvailable:        afterUpdate.IsAvailable,
	}

	// Confirmation on code
	return c.JSON(fiber.Map{
		"data":             codeResult,
		"status":           fiber.StatusOK,
		"error":      		result.Error,
		"rows_affected":    result.RowsAffected,
	})
}

func DeleteCode(c *fiber.Ctx) error {
	codeID := c.Params("codeID")

	var code models.Code
	database.DB.First(&code, codeID)

	// Delete an existing record
	database.DB.Delete(&code)

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})

}