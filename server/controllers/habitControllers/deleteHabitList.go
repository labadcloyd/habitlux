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


func DeleteHabitList(c *fiber.Ctx) error {
	//* auth middleware
	token := middlewares.AuthMiddleware(c)
	if token == nil {
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	u64, err := strconv.ParseUint(claims.Issuer, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	owner_id := uint(u64)

	//* data validation
	reqData := new(ReqDeleteHabitList)
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

	//* deleting the habitlist
	habitlist := models.HabitList{}
	habits := models.Habit{}
	if err := database.DB.Model(&habitlist).
		Where("Owner_ID = ?", owner_id).
		Where("Habit_Name = ?", reqData.Habit_Name).
		Delete(&habitlist).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
	}
	// deleting all habits associated to that habit list name
	if err := database.DB.Model(&habits).
		Where("Owner_ID = ?", owner_id).
		Where("Habit_Name = ?", reqData.Habit_Name).
		Group("Habit_Name, Date_Created").
		Delete(&habits).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	log.Println("Successfully deleted habit List and its habits")
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Successfully deleted habit List and its habits",
	})
}