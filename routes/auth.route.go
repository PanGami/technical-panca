package routes

import (
	"github.com/PanGami/technical-panca/controllers"
	"github.com/PanGami/technical-panca/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/register", controllers.SignUpUser)
	auth.Post("/login", controllers.SignInUser)
	auth.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
	auth.Get("/refresh", controllers.RefreshAccessToken)
}
