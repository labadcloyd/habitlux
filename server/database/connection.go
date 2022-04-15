package database

import (
	"habit-tracker/helpers"
	"habit-tracker/models"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbUrl := helpers.GoDotEnvVariable("MYSQL_URL")
	
	connection, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	log.Println("Successfully connected to database")

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Habit{})
}