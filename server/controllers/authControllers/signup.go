package controllers

import (
	"habit-tracker/database"
	"habit-tracker/helpers"
	"habit-tracker/models"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)


func Signup(c *fiber.Ctx) error {
	// data validation
	data := new(ReqSignUp)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := helpers.ValidateStruct(*data)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// hashing password and formatting data
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	user := models.User {
		Username: data.Username,
		Password: password,
	}

	// saving user
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
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

	log.Println("Successfully registered user")
	return c.JSON(user)
}