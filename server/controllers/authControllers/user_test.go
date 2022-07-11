package controllers

import (
	"habit-tracker/setup"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	// Initial Setup
	db, app := setup.MockSetupApp()
	app.Get("/api/user", func(c *fiber.Ctx) error { User(c, db); return nil })

	setup.SetupMockDB(db, t)
	defer setup.ClearMockDB(db, t)

	cookie, _, err := setup.SetupMockAccount(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a test account", err)
	}

	// generating a fake user token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(3)),
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 1, 0)),
	})
	if SecretKey == "" {
		SecretKey = os.Getenv("SECRET_KEY")
	}
	fakeToken, _ := claims.SignedString([]byte(SecretKey))
	fakeCookie := http.Cookie{
		Name:     "jwt",
		Value:    fakeToken,
		Expires:  time.Now().AddDate(0, 1, 0),
		HttpOnly: true,
		SameSite: 1,
		Secure:   true,
	}

	// Test Cases
	tests := []struct {
		description  string       // description of the test case
		cookie       *http.Cookie // cookie that contains the token
		expectedCode int          // expected HTTP status code
		expectedBody string       // expected result data
	}{
		{
			description:  "Returns error with empty cookie",
			cookie:       nil,
			expectedCode: 401,
			expectedBody: "{\"message\":\"Unautherized\"}",
		},
		{
			description:  "Returns error when user does not exist",
			cookie:       &fakeCookie,
			expectedCode: 400,
			expectedBody: "{\"message\":\"No user found\"}",
		},
		{
			description:  "Returns success with cookie",
			cookie:       cookie,
			expectedCode: 200,
			expectedBody: "{\"id\":1,\"username\":\"demo\"}",
		},
	}

	for _, test := range tests {
		// setting the request
		req := httptest.NewRequest("GET", "/api/user", nil)
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
