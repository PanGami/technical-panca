package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/PanGami/technical-panca/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func CheckPostgresConnection(c *fiber.Ctx) error {
	db := initializers.DB
	err := db.Exec("SELECT 1").Error
	if err != nil {
		log.Printf("Failed to ping PostgreSQL: %s\n", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to ping PostgreSQL",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "✅ Postgres connected successfully...",
	})
}

func CheckRedisConnection(c *fiber.Ctx) error {
	ctx := context.TODO()
	value, err := initializers.RedisClient.Get(ctx, "test").Result()
	if err == redis.Nil {
		fmt.Println("key: test does not exist")
	} else if err != nil {
		log.Printf("Failed to get value from Redis: %s\n", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": value,
	})
}
