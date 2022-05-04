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
	habitNames := []models.HabitList{}
	// getting the list of habit names
	if err := database.DB.Model(&models.HabitList{}).
		Where("Owner_ID = ?", owner_id).
		Find(&habitNames).Error; err != nil {
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
		Order("Date_Created desc").
		Find(&habits).Error; err != nil {
			log.Println(err)
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err,
			})
	}

	//* formatting data
	habitListSummary := make([][]models.Habit, len(habitNames))
	// initializing slice size (creates a joker element in order to append elements later)
	for i := 0; i < len(habitNames); i++ {
		habitListSummary[i] = make([]models.Habit, 1, len(habits))
	}
	// appending elements
	for i := 0; i < len(habits); i++ {
		for j := 0; j < len(habitNames); j++ {
			if habitNames[j].Habit_Name == habits[i].Habit_Name {
				habitListSummary[j] = append(habitListSummary[j], habits[i])
				break
			}
		}
	}
	// removing first joker element in each slice
	for i := 0; i < len(habitNames); i++ {
		habitListSummary[i] = habitListSummary[i][1:]
	}

	habitFormatted := &ResGetUserHabits {
		HabitNames: habitNames,
		Habits: habitListSummary,
	}

	log.Println(habitFormatted)
	return c.JSON(habitFormatted)
}
