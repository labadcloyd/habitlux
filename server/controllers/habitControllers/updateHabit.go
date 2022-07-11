package controllers

import (
	"database/sql"
	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UpdateHabit(c *fiber.Ctx, db *sql.DB) error {
	//* auth middleware
	token, owner_id, err := middlewares.AuthMiddleware(c)
	if token == nil || owner_id == 0 || err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unautherized",
		})
	}

	//* data validation
	reqData := new(ReqUpdateHabit)
	if err = middlewares.BodyValidation(reqData, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	//* checking if habit with the same date exists
	duplicateHabit := models.Habit{}
	row := db.
		QueryRow(`
			SELECT id FROM habits 
			WHERE owner_id = $1 AND date_created = $2 AND habit_list_id = $3`,
			owner_id, reqData.Date_Created, reqData.Habit_List_ID,
		)
	err = row.Scan(&duplicateHabit.ID)
	// only checking if the error is not caused by empty rows
	if err != nil && err != sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// returning error if record exists
	if (duplicateHabit != models.Habit{}) {
		if duplicateHabit.ID != 0 && duplicateHabit.ID != reqData.ID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Habit already exists",
			})
		}
	}

	//* updating the habit
	habit := models.Habit{
		ID:                  reqData.ID,
		Habit_List_ID:       reqData.Habit_List_ID,
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
			date_created = $1, comment = $2, target_repeat_count = $3, repeat_count = $4
			WHERE owner_id = $5 AND id = $6`,
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
