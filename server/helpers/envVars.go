package helpers

import (
	"os"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalln("Error loading .env file in helper")
	// }

	return os.Getenv(key)
}
