package routes

import (
	"github.com/PanGami/technical-panca/controllers"
	"github.com/gofiber/fiber/v2"
)

func ConnRoute(router fiber.Router) {
	conn := router.Group("/check")
	conn.Get("/redis", controllers.CheckRedisConnection)
	conn.Get("/postgres", controllers.CheckPostgresConnection)
}
