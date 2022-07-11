package controllers

import (
	"database/sql"
	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func User(c *fiber.Ctx, db *sql.DB) error {
	token, owner_id, err := middlewares.AuthMiddleware(c)
	if token == nil || owner_id == 0 || err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unautherized",
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	user := models.User{}

	row := db.
		QueryRow("SELECT username, id FROM users WHERE id = $1", claims.Issuer)
	// scanning and returning error
	err = row.Scan(&user.Username, &user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "No user found",
			})
		}
		log.Println("Error: ", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "An error occured in scanning user",
		})
	}
	return c.JSON(user)
}
