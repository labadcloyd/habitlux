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
	newHabitList := models.HabitList{
		ID:                   reqData.ID,
		Owner_ID:             owner_id,
		Habit_Name:           reqData.Habit_Name,
		Icon_Url:             reqData.Icon_Url,
		Color:                reqData.Color,
		Default_Repeat_Count: reqData.Default_Repeat_Count,
	}
	// updating habit list name
	if _, err := database.DB.
		Exec(`UPDATE habit_lists
			SET
				habit_name = $1, icon_url = $2, color = $3, default_repeat_count = $4
			WHERE owner_id = $5 AND id = $6`,
			reqData.Habit_Name, reqData.Icon_Url, reqData.Color, reqData.Default_Repeat_Count,
			owner_id, reqData.ID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Println("Successfully updated habit List and its habits")
	return c.Status(fiber.StatusOK).JSON(newHabitList)
}
