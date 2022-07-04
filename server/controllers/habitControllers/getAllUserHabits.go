package controllers

import (
	"database/sql"
	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllUserHabits(c *fiber.Ctx, db *sql.DB) error {
	//* auth middleware
	token, owner_id, err := middlewares.AuthMiddleware(c)
	if token == nil || owner_id == 0 || err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unautherized",
		})
	}

	//* data validation
	reqData := new(ReqGetUserHabits)
	if err = middlewares.BodyValidation(reqData, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//* querying the data
	habitListMap := make(map[int]int)
	habitListFormatted := make([]ResGetUserHabits, 0, 100)
	// getting the list of habit names
	rows, err := db.
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
	rows2, err := db.
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
