package routes

import (
	"github.com/PanGami/technical-panca/controllers"
	"github.com/PanGami/technical-panca/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Get("/me", middleware.DeserializeUser, controllers.GetMe)
}
