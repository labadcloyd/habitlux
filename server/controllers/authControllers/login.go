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


func Login(c *fiber.Ctx) error {
	// data validation
	reqData := new(ReqLogin)
	if err := c.BodyParser(&reqData); err != nil {
		return err
	}
	errors := helpers.ValidateStruct(*reqData)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var user = models.User{}

	// checking if user exists
	row, err := database.DB.
		Query("SELECT * FROM users WHERE username = $1 LIMIT 1", reqData.Username); 
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map {
			"message": "An error occured in scanning user",
		})
	}
	defer row.Close()
	// scanning and returning error
	if err := row.Scan(&user.Username, &user.ID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// checking if password matches user
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(reqData.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "inccorect password",
		})
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

	log.Println("Successfully logged user in")
	return c.JSON(fiber.Map{
		"message": "Successfully logged in",
	})
}