package controllers

import (
	"database/sql"
	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx, db *sql.DB) error {
	// data validation
	reqData := new(ReqSignUp)
	if err := middlewares.BodyValidation(reqData, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// hashing password and formatting reqData
	password, _ := bcrypt.GenerateFromPassword([]byte(reqData.Password), 10)
	user := models.User{
		Username: reqData.Username,
		Password: password,
	}

	// saving user
	row := db.
		QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", user.Username, user.Password)

	err := row.Scan(&user.ID)
	if err != nil {
		log.Println("Error: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	// generating jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
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
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().AddDate(0, 1, 0),
		HTTPOnly: true,
		SameSite: "None",
		Secure:   true,
	}

	c.Cookie(&cookie)

	log.Println("Successfully registered user")
	return c.JSON(user)
}
