package controllers

import (
	"time"

	"github.com/PanGami/technical-panca/initializers"
	"github.com/PanGami/technical-panca/middleware"
	"github.com/PanGami/technical-panca/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateProduct(c *fiber.Ctx) error {
	var payload *models.CreateProductInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := middleware.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})

	}

	newProduct := models.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Stock:       payload.Stock,
	}

	result := initializers.DB.Create(&newProduct)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": payload})
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	result := initializers.DB.Find(&products)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": products})
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	result := initializers.DB.First(&product, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Product not found"})
		} else {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": product})
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var existingProduct models.Product

	if err := initializers.DB.First(&existingProduct, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Product not found"})
		} else {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}
	}

	var payload models.UpdateProductInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := middleware.ValidateStruct(&payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	currentTime := time.Now()
	if payload.Name != nil {
		existingProduct.Name = *payload.Name
	}
	if payload.Description != nil {
		existingProduct.Description = *payload.Description
	}
	if payload.Price != nil {
		existingProduct.Price = *payload.Price
	}
	if payload.Stock != nil {
		existingProduct.Stock = *payload.Stock
	}
	existingProduct.UpdatedAt = &currentTime

	if err := initializers.DB.Save(&existingProduct).Error; err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": existingProduct})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	result := initializers.DB.First(&product, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Product not found"})
		} else {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
		}
	}

	result = initializers.DB.Delete(&product)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Product deleted successfully"})
}
