package middlewares

import (
	"habit-tracker/helpers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = helpers.GoDotEnvVariable("SECRET_KEY")

func AuthMiddleware(c *fiber.Ctx) *jwt.Token {
	cookie := c.Cookies("jwt")

	// parsing token
	token, err := jwt.ParseWithClaims(
		cookie, 
		&jwt.RegisteredClaims{}, 
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusUnauthorized)
		return nil
	}
	return token
}
		