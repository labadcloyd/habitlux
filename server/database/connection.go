package database

import (
	"habit-tracker/helpers"
	"log"
	"time"

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
	DB.SetConnMaxIdleTime(time.Minute * 2)
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")
}