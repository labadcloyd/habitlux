package setup

import (
	"database/sql"
	"habit-tracker/helpers"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var DB *sql.DB

func SetupApp(db *sql.DB) *fiber.App {
	DB = db
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			// Send custom error page
			err = ctx.Status(code).SendFile("./build/notfound.html")
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			// Return from handler
			return nil
		},
	})

	// allowing clients from different urls to access server
	// it is very important that we use the cors config first before-
	// declaring any routes
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	return app
}

func ConnectDB() *sql.DB {
	connStr := helpers.GoDotEnvVariable("POSTGRES_URL")
	var err error
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxIdleTime(time.Minute * 2)
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")
	return db
}
