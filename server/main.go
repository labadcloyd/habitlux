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
	app := fiber.New()

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
