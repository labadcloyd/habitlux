package controllers

import (
	"errors"
	"habit-tracker/database"
	"habit-tracker/helpers"
	"habit-tracker/middlewares"
	"habit-tracker/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
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

	if err := database.DB.Model(&models.Habit{}).
	Where("Owner_ID = ?", owner_id).
	Where("Date_Created = ?", reqData.Date_Created).
	First(&oldHabit).Error; err != nil {
		if ( !(errors.Is(err, gorm.ErrRecordNotFound)) ) {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	// updating habit if it already exists
	if (oldHabit != models.Habit{}) {
		habit := models.Habit {
			ID:										oldHabit.ID,
			Owner_ID: 						owner_id,
			Habit_Name:						reqData.Habit_Name,
			Date_Created: 				reqData.Date_Created,
			Comment: 							reqData.Comment,
			Target_Repeat_Count: 	reqData.Target_Repeat_Count,
			Repeat_Count: 				reqData.Repeat_Count,
		}
		if err := database.DB.Model(models.Habit{}).
			Where("Owner_ID = ?", owner_id).
			Where("ID = ?", oldHabit.ID).
			Updates(
				map[string]interface{}{
					"id": oldHabit.ID, 
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
	
		log.Println("Successfully updated habit")
		return c.Status(fiber.StatusOK).JSON(habit)
	}

	//* saving the habit
	habit := models.Habit {
		Owner_ID: 						owner_id,
		Habit_Name:						reqData.Habit_Name,
		Date_Created: 				reqData.Date_Created,
		Comment: 							reqData.Comment,
		Target_Repeat_Count: 	reqData.Target_Repeat_Count,
		Repeat_Count: 				reqData.Repeat_Count,
	}
	if err := database.DB.Create(&habit).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Println("Successfully saved habbit")
	return c.JSON(habit)
}