package controllers

import (
	"database/sql"
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

	row := database.DB.
		QueryRow("SEELCT * FROM users WHERE id = $1", claims.Issuer)
	// scanning and returning error
	err = row.Scan(&user.Username, &user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "No user found",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "An error occured in scanning user",
		})
	}
	return c.JSON(user)
}