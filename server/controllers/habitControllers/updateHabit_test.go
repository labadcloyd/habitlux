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

func TestUpdateHabit(t *testing.T) {
	// Initial Setup
	db, app := setup.MockSetupApp()
	app.Put("/api/habit", func(c *fiber.Ctx) error { UpdateHabit(c, db); return nil })

	setup.SetupMockDB(db, t)
	defer setup.ClearMockDB(db, t)

	cookie, owner_id, err := setup.SetupMockAccount(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a test account", err)
	}

	// Variables
	type testPayload struct {
		ID                  uint   `json:"id"`
		Habit_Name          string `json:"habit_name"`
		Date_Created        string `json:"date_created"`
		Comment             string `json:"comment"`
		Target_Repeat_Count uint   `json:"target_repeat_count"`
		Repeat_Count        uint   `json:"repeat_count"`
	}
	completePayload := testPayload{
		ID:                  1,
		Habit_Name:          "test2",
		Date_Created:        "2022-05-02",
		Comment:             "",
		Target_Repeat_Count: 4,
		Repeat_Count:        1,
	}
	incompletePayload := testPayload{
		Date_Created:        "2022-02-01",
		Comment:             "",
		Target_Repeat_Count: 4,
		Repeat_Count:        1,
	}
	incompletePayload2 := testPayload{
		Habit_Name:          "test",
		Comment:             "",
		Target_Repeat_Count: 4,
		Repeat_Count:        1,
	}
	violatingPayload := testPayload{
		ID:                  1,
		Habit_Name:          "asdasdasdasdasdasdasdasdasdasdasdasasdasdasddaasd",
		Date_Created:        "2022-02-01",
		Comment:             "",
		Target_Repeat_Count: 4,
		Repeat_Count:        1,
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
			description:  "Returns error with empty habit name",
			body:         incompletePayload,
			cookie:       cookie,
			expectedCode: 400,
			expectedBody: "{\"message\":\"failed validating data\"}",
			setupFunc:    func() {},
		}, {
			description:  "Returns error with empty date created field",
			body:         incompletePayload2,
			cookie:       cookie,
			expectedCode: 400,
			expectedBody: "{\"message\":\"failed validating data\"}",
			setupFunc:    func() {},
		}, {
			description:  "Returns error when habit name is too long",
			body:         violatingPayload,
			cookie:       cookie,
			expectedCode: 400,
			expectedBody: "{\"message\":\"failed validating data\"}",
			setupFunc:    func() {},
		}, {
			description:  "Returns success when fields are complete and user is authenticated",
			body:         completePayload,
			cookie:       cookie,
			expectedCode: 200,
			expectedBody: "{\"id\":1,\"owner_id\":1,\"habit_list_id\":0,\"habit_name\":\"test2\",\"date_created\":\"2022-05-02T00:00:00Z\",\"comment\":\"\",\"target_repeat_count\":4,\"repeat_count\":1}",
			setupFunc: func() {
				db.QueryRow(`
					INSERT INTO
					habit_lists (owner_id, habit_name, icon_url, color, default_repeat_count)
					VALUES ($1, $2, $3, $4, $5) RETURNING id`,
					owner_id, "test", "", "#ffffff", 2,
				)
				db.QueryRow(`
					INSERT INTO
					habits (owner_id, habit_name, habit_list_id, date_created, comment, target_repeat_count, repeat_count)
					VALUES ($1, $2, $3, $4, $5, $6, $7)
					RETURNING id`,
					owner_id, "test", 1, "2022-02-01", "", 1, 1,
				)
			},
		},
		// TODO Add this edge case to the endpoint
		// {
		// 	description:  "Returns error when habit with the same date_created exists",
		// 	body:         completePayload,
		// 	cookie:       cookie,
		// 	expectedCode: 400,
		// 	expectedBody: "{\"message\":\"Habit already exists\"}",
		// 	setupFunc:    func() {},
		// },
	}

	for _, test := range tests {
		// running optional setup function for test case
		test.setupFunc()

		// getting the body
		payload, _ := json.Marshal(test.body)

		// setting the request
		req := httptest.NewRequest("PUT", "/api/habit", bytes.NewBuffer(payload))
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
