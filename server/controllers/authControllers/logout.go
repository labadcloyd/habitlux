package controllers

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func Logout(c *fiber.Ctx) error {
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
