package controllers

import (
	"database/sql"

	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func CreateHabit(c *fiber.Ctx, db *sql.DB) error {
	//* auth middleware
	token, owner_id, err := middlewares.AuthMiddleware(c)
	if token == nil || owner_id == 0 || err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unautherized",
		})
	}

	//* data validation
	reqData := new(ReqCreateHabit)
	if err = middlewares.BodyValidation(reqData, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//* checking if record exists
	oldHabit := models.Habit{}
	row := db.
		QueryRow(`
			SELECT id FROM habits 
			WHERE owner_id = $1 AND date_created = $2 AND habit_list_id = $3`,
			owner_id, reqData.Date_Created, reqData.Habit_List_ID,
		)
	err = row.Scan(&oldHabit.ID)
	// only checking if the error is not caused by empty rows
	if err != nil && err != sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// returning error if record exists
	if (oldHabit != models.Habit{}) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Habit already exists",
		})
	}

	//* saving the habit
	habit := models.Habit{
		Owner_ID:            owner_id,
		Habit_Name:          reqData.Habit_Name,
		Habit_List_ID:       reqData.Habit_List_ID,
		Date_Created:        reqData.Date_Created,
		Comment:             reqData.Comment,
		Repeat_Count:        reqData.Repeat_Count,
		Target_Repeat_Count: reqData.Target_Repeat_Count,
	}
	row = db.QueryRow(`
		INSERT INTO
		habits (owner_id, habit_name, habit_list_id, date_created, comment, target_repeat_count, repeat_count)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`,
		owner_id,
		reqData.Habit_Name,
		reqData.Habit_List_ID,
		reqData.Date_Created,
		reqData.Comment,
		reqData.Target_Repeat_Count,
		reqData.Repeat_Count,
	)
	err = row.Scan(&habit.ID)
	if err, ok := err.(*pq.Error); ok {
		if err.Code.Name() == "foreign_key_violation" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": `Habit list: '` + reqData.Habit_Name + `' does not exist`,
			})
		} else {
			log.Println("Error: ", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Message,
			})
		}
	}

	log.Println("Successfully saved habbit")
	return c.JSON(habit)
}
