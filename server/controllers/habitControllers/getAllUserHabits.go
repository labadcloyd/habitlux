package controllers

import (
	"habit-tracker/database"
	"habit-tracker/helpers"
	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetAllUserHabits(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	owner_id := uint(u64)

	//* data validation
	reqData := ReqGetUserHabits{}
	if err := c.QueryParser(&reqData); err != nil {
		log.Println("err: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := helpers.ValidateStruct(reqData)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	//* querying the data
	habitListMap := make(map[int]int)
	habitListFormatted := make([]ResGetUserHabits, 0, 100)
	// getting the list of habit names
	rows, err := database.DB.
		Query(`SELECT * FROM habit_lists WHERE owner_id = $1`, owner_id)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		var newHabitList ResGetUserHabits
		newHabitList.Habits = make([]models.Habit, 0, 100)
		if err := rows.
			Scan(
				&newHabitList.ID,
				&newHabitList.Owner_ID,
				&newHabitList.Habit_Name,
				&newHabitList.Icon_Url,
				&newHabitList.Color,
				&newHabitList.Default_Repeat_Count,
			); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": err,
				})
		}
		habitListFormatted = append(habitListFormatted, newHabitList)
		habitListMap[int(newHabitList.ID)] = i
		i++
	}
	// getting habits
	rows2, err := database.DB.
		Query(`SELECT * FROM habits
		WHERE owner_ID = $1 AND date_Created BETWEEN $2 AND $3
		ORDER BY habit_Name, date_Created asc`,
		owner_id, reqData.Start_Date, reqData.End_Date)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}
	defer rows2.Close()
	for rows2.Next() {
		var newHabit models.Habit
		if err := rows2.
			Scan(
				&newHabit.ID,
				&newHabit.Owner_ID,
				&newHabit.Habit_List_ID,
				&newHabit.Habit_Name,
				&newHabit.Date_Created,
				&newHabit.Comment,
				&newHabit.Target_Repeat_Count,
				&newHabit.Repeat_Count,
			); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": err,
				})
		}
		habitListFormatted[habitListMap[int(newHabit.Habit_List_ID)]].Habits = 
			append(habitListFormatted[habitListMap[int(newHabit.Habit_List_ID)]].Habits, newHabit)
	}
	return c.JSON(habitListFormatted)
}
