package controllers

import (
	"habit-tracker/setup"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateHabit(t *testing.T) {
	// Initial Setup
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db, app := setup.MockSetupApp()
	app.Post("/habit", CreateHabit)

	setup.SetupMockDB(db, t)
	log.Fatal(app.Listen(":3001"))
	defer setup.ClearMockDB(db, t)

	cookie, err := setup.SetupMockAccount(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a test account", err)
	}

	// Test Cases
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{}

	for _, test := range tests {
		req := httptest.NewRequest("POST", test.route, nil)
		req.AddCookie(cookie)
		resp, _ := app.Test(req, -1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
