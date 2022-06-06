package controllers

import (
	"database/sql"
	"habit-tracker/database"
	"habit-tracker/models"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)


func DemoLogin(c *fiber.Ctx) error {
	// checking if user exists
	var user = models.User{}
	row := database.DB.QueryRow("SELECT * FROM users WHERE username = $1", "demo")
	// scanning and returning error
	err := row.Scan(&user.Username, &user.ID)
	if err == sql.ErrNoRows {
		// hashing password and formatting data
		password, _ := bcrypt.GenerateFromPassword([]byte("vErYSeCuRePaSsWoRd123!"), 10)
		user = models.User {
			Username: "demo",
			Password: password,
		}
		// saving user
		row2 := database.DB.
			QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", user.Username, user.Password)
		// scanning and returning error
		if err = row2.Scan(&user.ID); err != nil {
			log.Println("Error: ", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "An error occured in scanning user after query",
			})
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
		SameSite: "None",
		Secure: true,
	}

	c.Cookie(&cookie)

	log.Println("Successfully logged demo user in")
	return c.JSON(fiber.Map{
		"message": "Successfully logged in",
	})
}