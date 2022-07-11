package controllers

import (
	"bytes"
	"encoding/json"
	"habit-tracker/setup"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Initial Setup
	db, app := setup.MockSetupApp()
	app.Post("/api/login", func(c *fiber.Ctx) error { Login(c, db); return nil })

	setup.SetupMockDB(db, t)
	defer setup.ClearMockDB(db, t)

	_, _, err := setup.SetupMockAccount(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a test account", err)
	}

	// Variables
	completePayload := ReqLogin{
		Username: "demo",
		Password: "vErYSeCuRePaSsWoRd123!",
	}
	wrongUsername := ReqLogin{
		Username: "demo24",
		Password: "vErYSeCuRePaSsWoRd123!",
	}
	wrongPassword := ReqLogin{
		Username: "demo",
		Password: "vErYSeCu123",
	}
	// Test Cases
	tests := []struct {
		description  string                   // description of the test case
		body         interface{}              // request body data
		expectedCode int                      // expected HTTP status code
		expectedBody string                   // expected result data
		endingFunc   func(res *http.Response) // func to be executed before start
	}{
		{
			description:  "Returns error with empty body",
			body:         "",
			expectedCode: 400,
			expectedBody: "{\"message\":\"failed validating data\"}",
			endingFunc:   func(res *http.Response) {},
		}, {
			description:  "Returns error when username does not exist",
			body:         wrongUsername,
			expectedCode: 400,
			expectedBody: "{\"message\":\"user not found\"}",
			endingFunc:   func(res *http.Response) {},
		}, {
			description:  "Returns error with wrong password",
			body:         wrongPassword,
			expectedCode: 400,
			expectedBody: "{\"message\":\"inccorect password\"}",
			endingFunc:   func(res *http.Response) {},
		}, {
			description:  "Returns success when fields are complete",
			body:         completePayload,
			expectedCode: 200,
			expectedBody: "{\"message\":\"Successfully logged in\"}",
			endingFunc: func(res *http.Response) {
				// testing the cookie
				cookies := res.Cookies()
				for _, resCookie := range cookies {
					if resCookie.Name == "jwt" {
						_, err := jwt.ParseWithClaims(
							resCookie.Value,
							&jwt.RegisteredClaims{},
							func(t *jwt.Token) (interface{}, error) {
								return []byte(SecretKey), nil
							},
						)
						assert.Equalf(t, true, err == nil, "Returns success when fields are complete")
					}
				}
			},
		},
	}

	for _, test := range tests {
		// getting the body
		payload, _ := json.Marshal(test.body)

		// setting the request
		req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
		req.Header.Add("Accept", "application/json, text/plain, */*")
		req.Header.Add("Accept-Encoding", "gzip, deflate, br")
		req.Header.Add("Content-Type", "application/json")

		// testing the request
		res, _ := app.Test(req, -1)
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		bodyBytes, err := io.ReadAll(res.Body)
		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		bodyString := string(bodyBytes)
		assert.Equalf(t, test.expectedBody, bodyString, test.description)

		// running optional setup function for test case
		test.endingFunc(res)
	}
}
