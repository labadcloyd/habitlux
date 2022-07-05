package controllers

import (
	"database/sql"

	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func CreateHabitList(c *fiber.Ctx, db *sql.DB) error {
	//* auth middleware
	token, owner_id, err := middlewares.AuthMiddleware(c)
	if token == nil || owner_id == 0 || err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unautherized",
		})
	}

	//* data validation
	reqData := new(ReqCreateHabitList)
	if err = middlewares.BodyValidation(reqData, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//* checking if habitlist already exists
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

	//* saving the habitlist
	habit := models.HabitList{
		Owner_ID:             owner_id,
		Habit_Name:           reqData.Habit_Name,
		Color:                reqData.Color,
		Default_Repeat_Count: reqData.Default_Repeat_Count,
	}
	row = db.QueryRow(`
		INSERT INTO
		habit_lists (owner_id, habit_name, icon_url, color, default_repeat_count)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		owner_id, reqData.Habit_Name, reqData.Icon_Url, reqData.Color, reqData.Default_Repeat_Count,
	)
	err = row.Scan(&habit.ID)
	if err, ok := err.(*pq.Error); ok {
		if err.Code.Name() == "unique_violation" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": `Habit list: '` + reqData.Habit_Name + `' already exists`,
			})
		} else {
			log.Println("Error: ", err, "Error code: ", err.Code.Name())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Message,
			})
		}
	}

	return c.JSON(habit)
}
