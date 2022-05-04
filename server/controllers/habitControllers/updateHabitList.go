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


func UpdateHabitList(c *fiber.Ctx) error {
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
	reqData := new(ReqUpdateHabitList)
	if err := c.BodyParser(&reqData); err != nil {
		log.Println("err: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := helpers.ValidateStruct(*reqData)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	//* updating the habitList
	// getting the old habit list first
	oldHabitList := models.HabitList{}
	if err := database.DB.Model(&models.HabitList{}).
		Where("Owner_ID = ?", owner_id).
		Where("ID = ?", reqData.ID).
		Find(&oldHabitList).Error; err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
	}
	newHabitList := models.HabitList {
		Owner_ID: 						owner_id,
		Icon_Url: 						reqData.Icon_Url,
		Color: 								reqData.Color,
		Default_Repeat_Count: reqData.Default_Repeat_Count,
	}
	// updating habit list name
	if err := database.DB.Model(&models.HabitList{}).
		Where("Owner_ID = ?", owner_id).
		Where("ID = ?", reqData.ID).
		Updates(newHabitList).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
	}
	// updating all habits to the new habit list name
	if err := database.DB.Model(&models.Habit{}).
		Where("Owner_ID = ?", owner_id).
		Where("Habit_Name = ?", oldHabitList.Habit_Name).
		Select("Habit_Name").
		Updates(models.Habit{Habit_Name: reqData.Habit_Name}).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	log.Println("Successfully updated habit List and its habits")
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Successfully updated habit List and its habits",
	})
}