package helpers

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// validator function
type ErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var Validate = validator.New()

func ValidatePassword(fl validator.FieldLevel) bool {
	upperCase, _ := regexp.Compile("[A-Z]")
	lowerCase, _ := regexp.Compile("[a-z]")
	numbers, _ := regexp.Compile("[0-9]+")
	if upperCase.MatchString(fl.Field().String()) {
		if lowerCase.MatchString(fl.Field().String()) {
			if numbers.MatchString(fl.Field().String()) {
				return true
			}
		}
	}
	return false
}

func ValidateStruct(data interface{}) []*ErrorResponse {
	Validate.RegisterValidation("missingRequiredCharacters", ValidatePassword)
	var errors []*ErrorResponse
	err := Validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
