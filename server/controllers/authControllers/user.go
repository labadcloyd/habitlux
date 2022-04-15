package controllers

import (
	"habit-tracker/database"
	"habit-tracker/models"
	"log"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2"
)

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	// parsing token
	token, err := jwt.ParseWithClaims(
		cookie, 
		&jwt.StandardClaims{}, 
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
	
	claims := token.Claims.(*jwt.StandardClaims)

	user := models.User{}

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error;
		err != nil {
			log.Println(err)
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "No user found",
			})
		}

	return c.JSON(user)
}