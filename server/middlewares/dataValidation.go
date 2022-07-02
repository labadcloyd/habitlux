package middlewares

import (
	"errors"
	"habit-tracker/helpers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func BodyValidation(ReqStruct interface{}, c *fiber.Ctx) error {
	if err := c.BodyParser(ReqStruct); err != nil {
		log.Println("err: ", err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}
	err := helpers.ValidateStruct(&ReqStruct)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
		return errors.New("failed validating data")
	}
	return nil
}
func QueryValidation(ReqStruct interface{}, c *fiber.Ctx) error {
	if err := c.QueryParser(ReqStruct); err != nil {
		log.Println("err: ", err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}
	err := helpers.ValidateStruct(&ReqStruct)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
		return errors.New("failed validating data")
	}
	return nil
}
