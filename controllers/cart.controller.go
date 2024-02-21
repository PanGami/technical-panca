package controllers

import (
	"github.com/PanGami/technical-panca/initializers"
	"github.com/PanGami/technical-panca/middleware"
	"github.com/PanGami/technical-panca/models"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
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

	totalPrice := 0.0
	for _, item := range cartItems {
		totalPrice += calculateTotalPrice(item.ProductID, item.Quantity)
	}

	response := fiber.Map{
		"status":     "success",
		"data":       cartItems,
		"totalPrice": totalPrice,
	}

	return c.Status(fiber.StatusOK).JSON(response)
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

	var payload *models.CartCheckoutInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid request payload"})
	}

	var cartItems []models.Cart

	result := initializers.DB.Where("user_id = ?", payload.UserID).Find(&cartItems)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
	}

	totalPrice := 0.0
	for _, item := range cartItems {
		totalPrice += calculateTotalPrice(item.ProductID, item.Quantity)
	}

	if float64(payload.Money) < totalPrice {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Insufficient funds"})
	}

	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, item := range cartItems {
		var product models.Product
		pResult := tx.First(&product, "id = ?", item.ProductID)
		if pResult.Error != nil {
			tx.Rollback()
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": pResult.Error.Error()})
		}

		if product.Stock < item.Quantity {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Insufficient stock"})
		}

		product.Stock -= item.Quantity
		pUpdateResult := tx.Save(&product)
		if pUpdateResult.Error != nil {
			tx.Rollback()
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": pUpdateResult.Error.Error()})
		}
	}

	tx.Commit()

	// Clear the cart after payment successful
	if err := clearCart(payload.UserID); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Checkout successful"})
}

// Helper Functions
func calculateTotalPrice(productID *uuid.UUID, quantity int) float64 {
	var product models.Product

	result := initializers.DB.First(&product, "id = ?", productID)
	if result.Error != nil {
		// Handle error if product is not found
		return 0.0
	}

	return float64(quantity) * product.Price
}

func clearCart(userID string) error {
	result := initializers.DB.Delete(&models.Cart{}, "user_id = ?", userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
