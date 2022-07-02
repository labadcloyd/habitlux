package middlewares

import (
	"habit-tracker/helpers"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = helpers.GoDotEnvVariable("SECRET_KEY")

func AuthMiddleware(c *fiber.Ctx) (*jwt.Token, uint, error) {
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
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
		return nil, 0, err
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	u64, err := strconv.ParseUint(claims.Issuer, 10, 32)
	if err != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil, 0, err
	}

	owner_id := uint(u64)
	return token, owner_id, nil
}
