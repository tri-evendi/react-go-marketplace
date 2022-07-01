package controllers

import (
	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

type ProductDetails struct {
	ID           uint
	Name         string
	Description  string
	CategoryId   uint
	CategoryName string
	Price        float64
	ImagePath    string
}

func GetAllProducts(c *fiber.Ctx) error {
	var products = []models.Product{}
	database.DB.Preload("Category").Find(&products)

	return c.JSON(fiber.Map{
		"data":   products,
		"status": fiber.StatusOK,
	})
}

func ShowProduct(c *fiber.Ctx) error {
	productID := c.Params("productID")

	var product models.Product
	database.DB.Preload("Category").First(&product, productID)

	return c.JSON(fiber.Map{
		"data":   product,
		"status": fiber.StatusOK,
	})
}

func validCategoryIDs(categoryId_ uint) (bool, string) {
	valid := true
	var _ error
	var exists bool
	var allerror []bool
	var errorstring string

	var names = [4]string{"CategoryId"}

	_ = database.DB.Model(&models.Category{}).Select("count(*) > 0").Where("id", categoryId_).Find(&exists).Error
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

func StoreProduct(c *fiber.Ctx) error {
	data := new(models.Product)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Product failed.")
	}

	categoryId := data.CategoryID
	nameProduct := data.Name
	descriptionProduct := data.Description
	priceProduct := data.Price
	ImagePath := data.ImagePath

	categoryId_ := uint(categoryId)

	// Check IDs are vaild
	validity, screwedIds := validCategoryIDs(categoryId_)
	if !validity {
		return fiber.NewError(fiber.StatusNotFound, screwedIds+" doesn't exist")
	}

	product := models.Product{
		CategoryID:  categoryId_,
		Name:        nameProduct,
		Description: descriptionProduct,
		Price:       priceProduct,
		ImagePath:   ImagePath,
	}

	result := database.DB.Create(&product)
	database.DB.Preload("Category").Where("category_id = ?", product.CategoryID).Find(&product)

	productResult := ProductDetails{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		CategoryId:   product.CategoryID,
		CategoryName: product.Category.Name,
		Price:        product.Price,
		ImagePath:    product.ImagePath,
	}

	return c.JSON(fiber.Map{
		"data":          productResult,
		"status":        fiber.StatusOK,
		"error":         result.Error,
		"rows_affected": result.RowsAffected})
}

func UpdateProduct(c *fiber.Ctx) error {
	data := new(models.Product)
	err := c.BodyParser(&data)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Payload parser failed. Product failed.")
	}

	productID := c.Params("productID")

	categoryId := data.CategoryID
	nameProduct := data.Name
	descriptionProduct := data.Description
	priceProduct := data.Price
	ImagePath := data.ImagePath

	categoryId_ := uint(categoryId)

	// Check IDs are vaild
	validity, screwedIds := validCategoryIDs(categoryId_)
	if !validity {
		return fiber.NewError(fiber.StatusNotFound, screwedIds+" doesn't exist")
	}

	newProduct := models.Product{
		CategoryID:  categoryId_,
		Name:        nameProduct,
		Description: descriptionProduct,
		Price:       priceProduct,
		ImagePath:   ImagePath,
	}

	result := database.DB.Where("ID", productID).Updates(&newProduct)

	var afterUpdate models.Product
	database.DB.Preload("Category").First(&afterUpdate, productID)

	productResult := ProductDetails{
		ID:           afterUpdate.ID,
		Name:         afterUpdate.Name,
		Description:  afterUpdate.Description,
		CategoryId:   afterUpdate.CategoryID,
		CategoryName: afterUpdate.Category.Name,
		Price:        afterUpdate.Price,
		ImagePath:    afterUpdate.ImagePath,
	}

	return c.JSON(fiber.Map{
		"data":          productResult,
		"status":        fiber.StatusOK,
		"error":         result.Error,
		"rows_affected": result.RowsAffected,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("productID")

	var product models.Product
	database.DB.First(&product, productID)

	// Delete an existing record
	database.DB.Delete(&product)

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})

}
