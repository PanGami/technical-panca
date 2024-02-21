package routes

import (
	"github.com/PanGami/technical-panca/controllers"
	"github.com/gofiber/fiber/v2"
)

func CartRoutes(router fiber.Router) {
	cart := router.Group("/cart")
	cart.Post("/add", controllers.AddToCart)
	cart.Get("/items", controllers.GetCartItems)
	cart.Delete("/delete/:id", controllers.DeleteCartItem)
	cart.Post("/checkout", controllers.Checkout)
}
