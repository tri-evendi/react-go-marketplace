package controllers

import (
	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

type CategoryDetails struct {
	ID          uint
	Name        string
	Description string
}

func GetAllCategories(c *fiber.Ctx) error {
	var categories = []models.Category{}
	database.DB.Find(&categories)

	return c.JSON(fiber.Map{
		"data":   categories,
		"status": fiber.StatusOK,
	})
}

func GetProductByCategory(c *fiber.Ctx) error {
	categoryID := c.Params("categoryID")
	
	var products = []models.Product{}
	database.DB.Preload("Category").Where("category_id", categoryID).Find(&products)

	return c.JSON(fiber.Map{
		"data":   products,
		"status": fiber.StatusOK,
	})
}

func ShowCategory(c *fiber.Ctx) error {
	categoryID := c.Params("categoryID")
	
	var category models.Category
	database.DB.First(&category, categoryID)

	return c.JSON(fiber.Map{
		"data":   category,
		"status": fiber.StatusOK,
	})
}


func StoreCategory(c *fiber.Ctx) error {
	data := new(models.Category)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Category failed.")
	}

	nameCategory 	:= data.Name
	descriptionCategory 	:= data.Description

	category := models.Category{
		Name:     nameCategory,
		Description: descriptionCategory,
	}

	result := database.DB.Create(&category)

	database.DB.Where("name", category.Name).Find(&category)

	categoryResult := CategoryDetails{
		ID:      			category.ID,
		Name:   			category.Name,
		Description: 		category.Description,
	}

	return c.JSON(fiber.Map{
		"data":               categoryResult,
		"status":             fiber.StatusOK,
		"error":      result.Error,
		"rows_affected":      result.RowsAffected})
}

func UpdateCategory(c *fiber.Ctx) error {
	data := new(models.Category)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Category failed.")
	}

	categoryID := c.Params("categoryID")

	nameCategory 	:= data.Name
	descriptionCategory 	:= data.Name


	newCategory := models.Category{
		Name:     nameCategory,
		Description: descriptionCategory,
	}

	result := database.DB.Where("ID", categoryID).Updates(&newCategory)

	var afterUpdate models.Category
	database.DB.First(&afterUpdate, categoryID)

	categoryResult := CategoryDetails{
		ID:      			afterUpdate.ID,
		Name:   			afterUpdate.Name,
		Description: 		afterUpdate.Description,
	}

	return c.JSON(fiber.Map{
		"data":             categoryResult,
		"status":           fiber.StatusOK,
		"error":      		result.Error,
		"rows_affected":    result.RowsAffected,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	categoryID := c.Params("categoryID")

	var category models.Category
	database.DB.First(&category, categoryID)

	// Delete an existing record
	database.DB.Delete(&category)

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})

}