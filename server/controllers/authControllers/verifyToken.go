package controllers

import (
	"database/sql"
	"habit-tracker/middlewares"

	"github.com/gofiber/fiber/v2"
)

func VerifyToken(c *fiber.Ctx, db *sql.DB) error {
	//* auth middleware
	token, _, _ := middlewares.AuthMiddleware(c)
	if token == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Authenticated",
	})
}
