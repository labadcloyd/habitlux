package controllers

import "habit-tracker/helpers"

var SecretKey = helpers.GoDotEnvVariable("SECRET_KEY")

type ReqSignUp struct {
	Username    string  `json:"username" validate:"required,min=3,max=32"`
	Password		string 	`json:"password" validate:"required,min=10,max=32,missingRequiredCharacters"`
}
type ReqLogin struct {
	Username    string  `json:"username" validate:"required,min=6,max=32"`
	Password		string 	`json:"password" validate:"required,min=10,max=32,missingRequiredCharacters"`
}
