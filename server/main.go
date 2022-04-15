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

    routes.AuthRoutes(app)
    routes.HabitRoutes(app)

    // allowing clients from different urls to access server
    app.Use(cors.New(cors.Config{
        AllowCredentials: true,
    }))

    // returning 404 after wrong route
	app.Use(func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusNotFound).SendString("Error 404: not found")
    })

    log.Fatal(app.Listen(":3000"))
}