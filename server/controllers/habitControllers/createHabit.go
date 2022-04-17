package controllers

import (
	"habit-tracker/database"
	"habit-tracker/helpers"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
)


func CreateHabit(c *fiber.Ctx) error {
	// data validation
	data := new(ReqHabit)
	if err := c.BodyParser(&data); err != nil {
		log.Println("err: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := helpers.ValidateStruct(*data)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// saving the habit
	habit := models.Habit {
		Owner_ID: data.Owner_ID,
		Date_Created: data.Date_Created,
		Habit_Name: data.Habit_Name,
		Comment: data.Comment,
		Target_Repeat_Count: data.Target_Repeat_Count,
		Repeat_Count: data.Repeat_Count,
	}
	if err := database.DB.Create(&habit).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Println("Successfully saved habbit")
	return c.JSON(habit)
}