package controllers

import (
	"habit-tracker/helpers"
	"habit-tracker/middlewares"
	"habit-tracker/models"
	"habit-tracker/setup"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UpdateHabit(c *fiber.Ctx) error {
	db := setup.DB

	//* auth middleware
	token, owner_id, err := middlewares.AuthMiddleware(c)
	if token == nil || owner_id == 0 || err != nil {
		return err
	}

	//* data validation
	reqData := ReqUpdateHabit{}
	if err := c.BodyParser(&reqData); err != nil {
		log.Println("err: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := helpers.ValidateStruct(reqData)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	//* updating the habit
	habit := models.Habit{
		ID:                  reqData.ID,
		Owner_ID:            owner_id,
		Habit_Name:          reqData.Habit_Name,
		Date_Created:        reqData.Date_Created,
		Comment:             reqData.Comment,
		Target_Repeat_Count: reqData.Target_Repeat_Count,
		Repeat_Count:        reqData.Repeat_Count,
	}
	if _, err := db.
		Exec(`UPDATE habits
			SET
				habit_name = $1, date_created = $2, comment = $3, target_repeat_count = $4, repeat_count = $5
			WHERE owner_id = $6 AND id = $7`,
			reqData.Habit_Name,
			reqData.Date_Created,
			reqData.Comment,
			reqData.Target_Repeat_Count,
			reqData.Repeat_Count,
			owner_id, reqData.ID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Println("Successfully updated habbit")
	return c.Status(fiber.StatusOK).JSON(habit)
}
