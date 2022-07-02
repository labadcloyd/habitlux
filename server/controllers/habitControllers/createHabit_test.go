package controllers

import (
	"bytes"
	"encoding/json"
	"habit-tracker/setup"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type RouteRes struct {
}

func TestCreateHabit(t *testing.T) {
	// Initial Setup
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db, app := setup.MockSetupApp()
	app.Post("/api/habit", CreateHabit)

	setup.SetupMockDB(db, t)
	log.Fatal(app.Listen(":3001"))
	defer setup.ClearMockDB(db, t)

	cookie, err := setup.SetupMockAccount(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a test account", err)
	}

	// Test Cases
	tests := []struct {
		description  string       // description of the test case
		body         interface{}  // request body data
		cookie       *http.Cookie // cookie that contains the token
		expectedCode int          // expected HTTP status code
		expectedRes  interface{}  // expected result data
	}{
		{
			description:  "Returns error with empty cookie",
			body:         nil,
			cookie:       nil,
			expectedCode: 401,
			expectedRes:  RouteRes{},
		},
	}

	for _, test := range tests {
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(test.body)
		req := httptest.NewRequest("POST", "/api/habit", payloadBuf)
		req.AddCookie(cookie)
		resp, _ := app.Test(req, -1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		assert.Equalf(t, test.expectedRes, resp.Body, test.description)
	}
}
