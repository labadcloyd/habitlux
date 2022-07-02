package controllers

import (
	"habit-tracker/middlewares"
	"habit-tracker/setup"
	"log"

	"github.com/gofiber/fiber/v2"
)

func DeleteHabit(c *fiber.Ctx) error {
	db := setup.DB

	//* auth middleware
	token, owner_id, err := middlewares.AuthMiddleware(c)
	if token == nil || owner_id == 0 || err != nil {
		return err
	}

	//* data validation
	reqData := new(ReqDeleteHabit)
	if err = middlewares.BodyValidation(reqData, c); err != nil {
		return err
	}

	//* deleting the habit
	if _, err := db.Exec(`
		DELETE FROM habits WHERE owner_id = $1 AND id = $2`, owner_id, reqData.ID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Println("Successfully deleted habbit")
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Successfully deleted habbit",
	})
}
