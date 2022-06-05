package database

import (
	"habit-tracker/helpers"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := helpers.GoDotEnvVariable("POSTGRES_URL")
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")
}