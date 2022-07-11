package middlewares

import (
	"errors"
	"habit-tracker/helpers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func BodyValidation(ReqStruct interface{}, c *fiber.Ctx) error {
	if err := c.BodyParser(ReqStruct); err != nil {
		log.Println("err on line 13: ", err)
		return errors.New("failed validating data")
	}
	err := helpers.ValidateStruct(ReqStruct)
	if err != nil {
		log.Println("err on line 18: ", err)
		return errors.New("failed validating data")
	}
	return nil
}
func QueryValidation(ReqStruct interface{}, c *fiber.Ctx) error {
	if err := c.QueryParser(ReqStruct); err != nil {
		log.Println("err on line 25: ", err)
		return errors.New("failed validating data")
	}
	err := helpers.ValidateStruct(ReqStruct)
	if err != nil {
		log.Println("err on line 30: ", err)
		return errors.New("failed validating data")
	}
	return nil
}
