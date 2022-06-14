package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
