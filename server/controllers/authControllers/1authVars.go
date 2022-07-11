package controllers

import "os"

var SecretKey = os.Getenv("SECRET_KEY")

type ReqSignUp struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=10,max=32"` //,missingRequiredCharacters // add this for password validation
}
type ReqLogin struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=10,max=32"` //,missingRequiredCharacters  // add this for password validation
}
