package main

import (
	"habit-tracker/database"
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
	// also make sure to specify headers because it will block cookies if you dont
	app.Use(cors.New(cors.Config {
		AllowHeaders: "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		AllowCredentials: true,
	}))

	// routes
	routes.AuthRoutes(app)
	routes.HabitRoutes(app)


	// returning 404 after wrong route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Error 404: not found")
	})

	log.Fatal(app.Listen(":3001"))
}
