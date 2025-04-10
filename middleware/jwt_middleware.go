package middleware

import (
	"api_perpustakaan/utils"

	"github.com/gofiber/fiber/v2"
)

func Protected(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	_, err := utils.ValidateJWT(token)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
