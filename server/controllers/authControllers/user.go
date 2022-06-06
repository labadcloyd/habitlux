package controllers

import (
	"habit-tracker/database"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func User(c *fiber.Ctx) error {
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
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	
	claims := token.Claims.(*jwt.RegisteredClaims)

	user := models.User{}

	row, err := database.DB.
		Query("SEELCT * FROM users WHERE id = $1 LIMIT 1", claims.Issuer)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "No user found",
		})
	}
	defer row.Close()
	// scanning and returning error
	if err := row.Scan(&user.Username, &user.ID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "An error occured in scanning user",
		})
	}
	return c.JSON(user)
}