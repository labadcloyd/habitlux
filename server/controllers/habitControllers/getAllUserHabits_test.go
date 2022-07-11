package controllers

import (
	"encoding/json"
	"fmt"
	"habit-tracker/models"
	"habit-tracker/setup"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUserHabits(t *testing.T) {
	// Initial Setup
	db, app := setup.MockSetupApp()
	app.Get("/api/habit", func(c *fiber.Ctx) error { GetAllUserHabits(c, db); return nil })

	setup.SetupMockDB(db, t)
	defer setup.ClearMockDB(db, t)

	cookie, owner_id, err := setup.SetupMockAccount(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a test account", err)
	}

	// Setting up initial data
	habitListMap := make(map[int]int)
	habitListFormatted := make([]ResGetUserHabits, 0, 100)
	habitListEmptyHabits := make([]ResGetUserHabits, 0, 100)
	for i := 1; i < 4; i++ {
		hl := ResGetUserHabits{
			ID:                   uint(i),
			Owner_ID:             owner_id,
			Habit_Name:           fmt.Sprintf("test%v", i),
			Icon_Url:             "",
			Color:                "#ffffff",
			Default_Repeat_Count: 2,
		}
		hl.Habits = make([]models.Habit, 0, 100)
		db.QueryRow(`
			INSERT INTO
			habit_lists (owner_id, habit_name, icon_url, color, default_repeat_count)
			VALUES ($1, $2, $3, $4, $5)`,
			owner_id, fmt.Sprintf("test%v", i), "", "#ffffff", 2,
		)
		habitListFormatted = append(habitListFormatted, hl)
		habitListEmptyHabits = append(habitListEmptyHabits, hl)
		habitListMap[int(hl.ID)-1] = i - 1
		for j := 1; j < 4; j++ {
			row := db.QueryRow(`
				INSERT INTO
				habits (owner_id, habit_name, habit_list_id, date_created, comment, target_repeat_count, repeat_count)
				VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, date_created`,
				owner_id, fmt.Sprintf("test%v", i), i, fmt.Sprintf("2022-01-0%v", j), "", 1, 1,
			)
			h := models.Habit{
				Owner_ID:            owner_id,
				Habit_List_ID:       uint(i),
				Habit_Name:          fmt.Sprintf("test%v", i),
				Comment:             "",
				Target_Repeat_Count: 1,
				Repeat_Count:        1}
			err := row.Scan(&h.ID, &h.Date_Created)
			if err != nil {
				log.Println("ERROR ON LINE 51, GET ALL USER HABITS FILE: ", err)
			}
			habitListFormatted[habitListMap[int(h.Habit_List_ID)-1]].Habits =
				append(habitListFormatted[habitListMap[int(h.Habit_List_ID)-1]].Habits, h)
		}
	}

	jsonRes, _ := json.Marshal(habitListFormatted)
	jsonResEmpty, _ := json.Marshal(habitListEmptyHabits)

	// Test Cases
	tests := []struct {
		description  string            // description of the test case
		query        map[string]string // request query data
		cookie       *http.Cookie      // cookie that contains the token
		expectedCode int               // expected HTTP status code
		expectedBody string            // expected result data
	}{
		{
			description:  "Returns Unautherized error with empty cookie",
			query:        map[string]string{"": ""},
			cookie:       nil,
			expectedCode: 401,
			expectedBody: "{\"message\":\"Unautherized\"}",
		}, {
			description:  "Returns error with empty query",
			query:        map[string]string{"": ""},
			cookie:       cookie,
			expectedCode: 400,
			expectedBody: "{\"message\":\"failed validating data\"}",
		}, {
			description:  "Returns success with correct params",
			query:        map[string]string{"Start_Date": "2022-01-01", "End_Date": "2022-01-04"},
			cookie:       cookie,
			expectedCode: 200,
			expectedBody: string(jsonRes),
		}, {
			description:  "Returns habit lists with empty habits when dates given have no habits",
			query:        map[string]string{"Start_Date": "2022-02-01", "End_Date": "2022-02-04"},
			cookie:       cookie,
			expectedCode: 200,
			expectedBody: string(jsonResEmpty),
		},
	}

	for _, test := range tests {
		// setting the request
		url := fmt.Sprintf("/api/habit?start_date=%v&end_date=%v", test.query["Start_Date"], test.query["End_Date"])
		req := httptest.NewRequest("GET", url, nil)
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
		// Reading the response query should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		bodyString := string(bodyBytes)
		assert.Equalf(t, test.expectedBody, bodyString, test.description)
	}
}
