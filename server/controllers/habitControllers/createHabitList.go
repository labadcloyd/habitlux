package controllers

import (
	"habit-tracker/helpers"
	"habit-tracker/setup"

	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func CreateHabitList(c *fiber.Ctx) error {
	db := setup.DB

	//* auth middleware
	token, owner_id, err := middlewares.AuthMiddleware(c)
	if token == nil || owner_id == 0 || err != nil {
		return err
	}

	//* data validation
	reqData := new(ReqCreateHabitList)
	if err := c.BodyParser(&reqData); err != nil {
		log.Println("err: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	reqErrors := helpers.ValidateStruct(*reqData)
	if reqErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(reqErrors)
	}
	//* saving the habitlist
	habit := models.HabitList{
		Owner_ID:             owner_id,
		Habit_Name:           reqData.Habit_Name,
		Color:                reqData.Color,
		Default_Repeat_Count: reqData.Default_Repeat_Count,
	}
	row := db.QueryRow(`
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
