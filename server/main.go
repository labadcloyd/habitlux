package main

import (
	"habit-tracker/database"
	"habit-tracker/helpers"
	"habit-tracker/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
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
	app.Use(cors.New(cors.Config {
		AllowCredentials: true,
	}))
	
	// ui
	routes.StaticRoutes(app)

	// routes
	routes.AuthRoutes(app)
	routes.HabitRoutes(app)

	port := helpers.GoDotEnvVariable("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":"+ port))
}
