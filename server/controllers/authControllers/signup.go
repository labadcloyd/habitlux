package controllers

import (
	"habit-tracker/database"
	"habit-tracker/helpers"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
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
		Name: data.Name,
		Email: data.Email,
		Password: password,
	}

	// saving user
	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}

	log.Println("Successfully registered user")
	return c.JSON(user)
}