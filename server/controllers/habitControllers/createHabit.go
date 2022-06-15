package controllers

import (
	"database/sql"
	"habit-tracker/database"
	"habit-tracker/helpers"

	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lib/pq"
)

func CreateHabit(c *fiber.Ctx) error {
	//* auth middleware
	token := middlewares.AuthMiddleware(c)
	if token == nil {
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	u64, err := strconv.ParseUint(claims.Issuer, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	owner_id := uint(u64)

	//* data validation
	reqData := new(ReqCreateHabit)
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

	//* checking if record exists
	oldHabit := models.Habit{}
	row := database.DB.
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
	row = database.DB.QueryRow(`
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
