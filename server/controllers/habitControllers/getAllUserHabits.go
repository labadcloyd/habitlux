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
	habits := []models.Habit{}
	habitList := []models.HabitList{}
	habitListFormatted := []ResGetUserHabits{}
	// getting the list of habit names
	if err := database.DB.Model(&models.HabitList{}).
		Where("Owner_ID = ?", owner_id).
		Find(&habitList).Error; err != nil {
			log.Println(err)
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err,
			})
	}
	// getting habits
	if err := database.DB.Model(&models.Habit{}).
		Where("Owner_ID = ?", owner_id).
		Where("Date_Created BETWEEN ? AND ?", reqData.Start_Date, reqData.End_Date).
		Group("Habit_Name, Date_Created").
		Order("Date_Created asc").
		Find(&habits).Error; err != nil {
			log.Println(err)
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err,
			})
	}
	if len(habitList) < 1 {
		return c.JSON([]models.HabitList{})
	}

	//* formatting data
	// initializing slice size (creates a joker element in order to append elements later)
	for i := 0; i < len(habitList); i++ {
		newHabit := ResGetUserHabits {
			ID:										habitList[i].ID,
			Owner_ID:							habitList[i].Owner_ID,
			Habit_Name:						habitList[i].Habit_Name,
			Icon_Url: 						habitList[i].Icon_Url,	
			Color:								habitList[i].Color,
			Default_Repeat_Count: habitList[i].Default_Repeat_Count,
			Habits:								[]models.Habit{},
		}
		for j := 0; j < len(habits); j++ {
			if habitList[i].Habit_Name == habits[j].Habit_Name {
				newHabit.Habits = append(newHabit.Habits, habits[j])
				// * you can still Make this loop more efficient 
				// * by removing the elements that have already been appended
				// habits = append(habits[:j], habits[:]...)
			}
		}
		habitListFormatted = append(habitListFormatted, newHabit)
	}
	return c.JSON(habitListFormatted)
}
