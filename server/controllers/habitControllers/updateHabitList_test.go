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
	"github.com/stretchr/testify/assert"
)

func TestUpdateHabitList(t *testing.T) {
	// Initial Setup
	db, app := setup.MockSetupApp()
	app.Put("/api/habitlist", func(c *fiber.Ctx) error { UpdateHabitList(c, db); return nil })

	setup.SetupMockDB(db, t)
	defer setup.ClearMockDB(db, t)

	cookie, owner_id, err := setup.SetupMockAccount(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a test account", err)
	}

	// Variables
	type testPayload struct {
		ID                   uint   `json:"id" `
		Habit_Name           string `json:"habit_name"`
		Icon_Url             string `json:"icon_url"`
		Color                string `json:"color"`
		Default_Repeat_Count uint   `json:"default_repeat_count"`
	}
	completePayload := testPayload{
		ID:                   1,
		Habit_Name:           "test",
		Icon_Url:             "aasdasdasd",
		Color:                "#fff",
		Default_Repeat_Count: 1,
	}
	incompletePayload := testPayload{
		ID:                   1,
		Habit_Name:           "",
		Icon_Url:             "aasdasdasd",
		Color:                "#fff",
		Default_Repeat_Count: 1,
	}
	violatingPayload := testPayload{
		ID:                   1,
		Habit_Name:           "asdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasd",
		Icon_Url:             "aasdasdasd",
		Color:                "#fff",
		Default_Repeat_Count: 1,
	}

	// Test Cases
	tests := []struct {
		description  string       // description of the test case
		body         interface{}  // request body data
		cookie       *http.Cookie // cookie that contains the token
		expectedCode int          // expected HTTP status code
		expectedBody string       // expected result data
		setupFunc    func()       // func to be executed before start
	}{
		{
			description:  "Returns Unautherized error with empty cookie",
			body:         "",
			cookie:       nil,
			expectedCode: 401,
			expectedBody: "{\"message\":\"Unautherized\"}",
			setupFunc:    func() {},
		}, {
			description:  "Returns error with empty body",
			body:         "",
			cookie:       cookie,
			expectedCode: 400,
			expectedBody: "{\"message\":\"failed validating data\"}",
			setupFunc:    func() {},
		}, {
			description:  "Returns error when payload is incomplete",
			body:         incompletePayload,
			cookie:       cookie,
			expectedCode: 400,
			expectedBody: "{\"message\":\"failed validating data\"}",
			setupFunc:    func() {},
		}, {
			description:  "Returns error when payload violates a field",
			body:         violatingPayload,
			cookie:       cookie,
			expectedCode: 400,
			expectedBody: "{\"message\":\"failed validating data\"}",
			setupFunc:    func() {},
		}, {
			description:  "Returns success when authenticated and with complete payload",
			body:         completePayload,
			cookie:       cookie,
			expectedCode: 200,
			expectedBody: "{\"id\":1,\"owner_id\":1,\"habit_name\":\"test\",\"icon_url\":\"aasdasdasd\",\"color\":\"#fff\",\"default_repeat_count\":1}",
			setupFunc:    func() {},
		}, {
			description:  "Returns error when habit list name already exists",
			body:         completePayload,
			cookie:       cookie,
			expectedCode: 400,
			expectedBody: "{\"message\":\"Habit list already exists\"}",
			setupFunc: func() {
				db.QueryRow(`
					INSERT INTO
					habit_lists (owner_id, habit_name, icon_url, color, default_repeat_count)
					VALUES ($1, $2, $3, $4, $5) RETURNING id`,
					owner_id, "test", "", "#ffffff", 2,
				)
			},
		},
	}

	for _, test := range tests {
		// running optional setup function for test case
		test.setupFunc()

		// getting the body
		payload, _ := json.Marshal(test.body)

		// setting the request
		req := httptest.NewRequest("PUT", "/api/habitlist", bytes.NewBuffer(payload))
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
