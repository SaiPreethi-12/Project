package controllers

import (
	"encoding/json"
	"project/config"
	"project/models"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {

	cachedData, err := config.RedisClient.Get(config.Ctx, "users").Result()

	if err == nil {

		var cachedUsers []models.User

		json.Unmarshal([]byte(cachedData), &cachedUsers)

		return c.JSON(fiber.Map{
			"source": "redis cache",
			"data":   cachedUsers,
		})
	}

	var users []models.User

	config.DB.Find(&users)

	jsonData, _ := json.Marshal(users)

	config.RedisClient.Set(config.Ctx, "users", jsonData, 0)

	return c.JSON(fiber.Map{
		"source": "postgres database",
		"data":   users,
	})
}
