package controllers

import (
	"habit-tracker/setup"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestVerifyToken(t *testing.T) {
	// Initial Setup
	db, app := setup.MockSetupApp()
	app.Post("/api/verifytoken", func(c *fiber.Ctx) error { VerifyToken(c, db); return nil })

	setup.SetupMockDB(db, t)
	defer setup.ClearMockDB(db, t)

	cookie, _, err := setup.SetupMockAccount(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a test account", err)
	}

	// Test Cases
	tests := []struct {
		description  string       // description of the test case
		cookie       *http.Cookie // cookie that contains the token
		expectedCode int          // expected HTTP status code
		expectedBody string       // expected result data
	}{
		{
			description:  "Returns error when cookie does not exist",
			cookie:       nil,
			expectedCode: 400,
			expectedBody: "{\"message\":\"Unauthenticated\"}",
		},
		{
			description:  "Returns success and cookie",
			cookie:       cookie,
			expectedCode: 200,
			expectedBody: "{\"message\":\"Authenticated\"}",
		},
	}
	for _, test := range tests {
		// setting the request
		req := httptest.NewRequest("POST", "/api/verifytoken", nil)
		req.Header.Add("Accept", "application/json, text/plain, */*")
		req.Header.Add("Accept-Encoding", "gzip, deflate, br")
		req.Header.Add("Content-Type", "application/json")
		if test.cookie != nil {
			req.AddCookie(test.cookie)
		}

		// testing the request
		res, _ := app.Test(req, -1)
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		bodyBytes, err := io.ReadAll(res.Body)
		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		bodyString := string(bodyBytes)
		assert.Equalf(t, test.expectedBody, bodyString, test.description)
	}
}
