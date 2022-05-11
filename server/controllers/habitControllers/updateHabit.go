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


func UpdateHabit(c *fiber.Ctx) error {
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
	reqData := ReqUpdateHabit{}
	if err := c.BodyParser(&reqData); err != nil {
		log.Println("err: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := helpers.ValidateStruct(reqData)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	//* updating the habit
	habit := models.Habit {
		ID:										reqData.ID,
		Owner_ID: 						owner_id,
		Habit_Name:						reqData.Habit_Name,
		Date_Created: 				reqData.Date_Created,
		Comment: 							reqData.Comment,
		Target_Repeat_Count: 	reqData.Target_Repeat_Count,
		Repeat_Count: 				reqData.Repeat_Count,
	}
	if err := database.DB.Model(&habit).
		Where("Owner_ID = ?", owner_id).
		Where("ID = ?", reqData.ID).
		Updates(
			map[string]interface{}{
				"id": reqData.ID, 
				"owner_id": owner_id,
				"habit_name": reqData.Habit_Name,
				"date_created": reqData.Date_Created,
				"comment": reqData.Comment,
				"target_repeat_count": reqData.Target_Repeat_Count,
				"repeat_count": reqData.Repeat_Count,
			},
		).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	log.Println("Successfully updated habbit")
	return c.Status(fiber.StatusOK).JSON(habit)
}