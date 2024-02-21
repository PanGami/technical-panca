package routes

import (
	"github.com/PanGami/technical-panca/controllers"
	"github.com/PanGami/technical-panca/middleware"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router) {
	product := router.Group("/products")
	product.Post("/", middleware.DeserializeUser, controllers.CreateProduct)
	product.Get("/", middleware.DeserializeUser, controllers.GetProducts)
	product.Get("/category", middleware.DeserializeUser, controllers.GetProductsByCategory)
	product.Get("/:id", middleware.DeserializeUser, controllers.GetProduct)
	product.Put("/:id", middleware.DeserializeUser, controllers.UpdateProduct)
	product.Delete("/:id", middleware.DeserializeUser, controllers.DeleteProduct)
}
