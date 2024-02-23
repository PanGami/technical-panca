package routes

import (
	"github.com/PanGami/technical-panca/controllers"
	"github.com/PanGami/technical-panca/middleware"
	"github.com/gofiber/fiber/v2"
)

func CartRoutes(router fiber.Router) {
	cart := router.Group("/cart")
	cart.Post("/add", middleware.DeserializeUser, controllers.AddToCart)
	cart.Get("/items", middleware.DeserializeUser, controllers.GetCartItems)
	cart.Delete("/delete/:id", middleware.DeserializeUser, controllers.DeleteCartItem)
	cart.Post("/checkout", middleware.DeserializeUser, controllers.Checkout)
}
