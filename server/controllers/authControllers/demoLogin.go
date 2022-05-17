package controllers

import (
	"habit-tracker/database"
	"habit-tracker/models"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func DemoLogin(c *fiber.Ctx) error {
	// checking if user exists
	var user = models.User{}

	if err := database.DB.
		Where("username = ?", "demo").
		First(&user).Error; 
		err != nil {
			// hashing password and formatting data
			password, _ := bcrypt.GenerateFromPassword([]byte("vErYSeCuRePaSsWoRd123!"), 10)
			user = models.User {
				Username: "demo",
				Password: password,
			}
		
			// saving user
			if err := database.DB.Create(&user).Error; err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(err)
			}
	}
	
	// generating jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 1, 0)),
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	// saving jwt to cookie
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().AddDate(0, 1, 0),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	log.Println("Successfully logged demo user in")
	return c.JSON(fiber.Map{
		"message": "Successfully logged in",
	})
}