package controllers

import (
	"github.com/PanGami/technical-panca/initializers"
	"github.com/PanGami/technical-panca/middleware"
	"github.com/PanGami/technical-panca/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddToCart(c *fiber.Ctx) error {
	var payload *models.CartItemInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := middleware.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	NewItemCart := models.Cart{
		UserID:    payload.UserID,
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
	}

	result := initializers.DB.Create(&NewItemCart)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": payload})
}

func GetCartItems(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	var cartItems []models.Cart

	result := initializers.DB.Where("user_id = ?", userID).Find(&cartItems)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": cartItems})
}

func DeleteCartItem(c *fiber.Ctx) error {
	id := c.Params("id")

	var cartItem models.Cart

	result := initializers.DB.First(&cartItem, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Cart item not found"})
		} else {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
		}
	}

	result = initializers.DB.Delete(&cartItem)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Cart item deleted successfully"})
}

func Checkout(c *fiber.Ctx) error {
	// Implement your checkout logic here, e.g., processing payment, updating inventory, etc.
	// For simplicity, we'll just return a success message.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Checkout successful"})
}
