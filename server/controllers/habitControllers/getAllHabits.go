package controllers

import (
	"habit-tracker/database"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
)


func GetAllHabits(c *fiber.Ctx) error {
	habits := []models.Habit{}

	if err := database.DB.Find(&habits).Error;
		err != nil {
			log.Println(err)
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "No user found",
			})
		}
	return c.JSON(habits)
}