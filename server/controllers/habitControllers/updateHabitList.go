package controllers

import (
	"database/sql"
	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UpdateHabitList(c *fiber.Ctx, db *sql.DB) error {
	//* auth middleware
	token, owner_id, err := middlewares.AuthMiddleware(c)
	if token == nil || owner_id == 0 || err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unautherized",
		})
	}

	//* data validation
	reqData := new(ReqUpdateHabitList)
	if err = middlewares.BodyValidation(reqData, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//* checking if habitlist name already exists
	oldHabitList := models.HabitList{}
	row := db.
		QueryRow(`
			SELECT id FROM habit_lists 
			WHERE owner_id = $1 AND habit_name = $2`,
			owner_id, reqData.Habit_Name,
		)
	err = row.Scan(&oldHabitList.ID)
	// only checking if the error is not caused by empty rows
	if err != nil && err != sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// returning error if record exists
	if (oldHabitList != models.HabitList{}) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Habit list already exists",
		})
	}
	//* updating the habitList
	newHabitList := models.HabitList{
		ID:                   reqData.ID,
		Owner_ID:             owner_id,
		Habit_Name:           reqData.Habit_Name,
		Icon_Url:             reqData.Icon_Url,
		Color:                reqData.Color,
		Default_Repeat_Count: reqData.Default_Repeat_Count,
	}
	// updating habit list name
	if _, err := db.
		Exec(`UPDATE habit_lists
			SET
				habit_name = $1, icon_url = $2, color = $3, default_repeat_count = $4
			WHERE owner_id = $5 AND id = $6`,
			reqData.Habit_Name, reqData.Icon_Url, reqData.Color, reqData.Default_Repeat_Count,
			owner_id, reqData.ID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Println("Successfully updated habit List and its habits")
	return c.Status(fiber.StatusOK).JSON(newHabitList)
}
