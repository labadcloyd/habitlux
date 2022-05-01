package controllers

import (
	"log"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2"
)

func VerifyToken(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	// parsing token
	_, err := jwt.ParseWithClaims(
		cookie, 
		&jwt.RegisteredClaims{}, 
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Authenticated",
	})
}