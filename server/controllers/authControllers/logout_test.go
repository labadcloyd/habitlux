package controllers

import (
	"habit-tracker/setup"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogout(t *testing.T) {
	// Initial Setup
	db, app := setup.MockSetupApp()
	app.Post("/api/logout", func(c *fiber.Ctx) error { Logout(c, db); return nil })

	setup.SetupMockDB(db, t)
	defer setup.ClearMockDB(db, t)

	cookie, _, err := setup.SetupMockAccount(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a test account", err)
	}

	// Test Cases
	test := struct {
		description  string // description of the test case
		expectedCode int    // expected HTTP status code
		expectedBody string // expected result data
	}{
		description:  "Returns success and empty jwt cookie",
		expectedCode: 200,
		expectedBody: "{\"message\":\"Successfully logged out\"}",
	}
	// setting the request
	req := httptest.NewRequest("POST", "/api/logout", nil)
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(cookie)

	// testing the request
	res, _ := app.Test(req, -1)
	assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

	bodyBytes, err := io.ReadAll(res.Body)
	// Reading the response body should work everytime, such that
	// the err variable should be nil
	assert.Nilf(t, err, test.description)

	bodyString := string(bodyBytes)
	assert.Equalf(t, test.expectedBody, bodyString, test.description)

	// testing if the jwt cookie still exists
	cookieExists := true
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
			if err != nil {
				cookieExists = false
			}
		}
	}

	assert.Equalf(t, false, cookieExists, test.description)
}
