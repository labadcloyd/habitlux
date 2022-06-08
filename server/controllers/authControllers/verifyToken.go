package controllers

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/middlewares"
)

func VerifyToken(c *fiber.Ctx) error {
	//* auth middleware
	token := middlewares.AuthMiddleware(c)
	if token == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Authenticated",
	})
}