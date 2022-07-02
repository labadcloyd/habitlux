package controllers

import (
	"habit-tracker/helpers"
	"habit-tracker/middlewares"
	"habit-tracker/setup"
	"log"

	"github.com/gofiber/fiber/v2"
)

func DeleteHabitList(c *fiber.Ctx) error {
	db := setup.DB

	//* auth middleware
	token, owner_id, err := middlewares.AuthMiddleware(c)
	if token == nil || owner_id == 0 || err != nil {
		return err
	}

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
	if _, err := db.Exec(`
		DELETE FROM habit_lists WHERE owner_id = $1 AND id = $2`, owner_id, reqData.ID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Println("Successfully deleted habit List and its habits")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted habit List and its habits",
	})
}
