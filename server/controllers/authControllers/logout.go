package controllers

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx, db *sql.DB) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: "None",
		Secure:   true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
