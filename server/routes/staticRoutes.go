package routes

import (
	"github.com/gofiber/fiber/v2"
)

func StaticRoutes(app *fiber.App) {
	staticConfig := fiber.Static{
		Compress:      true,
		Browse:        true,
		Index:         "index.html",
	}
	app.Static("/", 					"./build", staticConfig)
	app.Static("/auth", 			"./build/auth.html", staticConfig)
	app.Static("/dashboard",	"./build/dashboard.html", staticConfig)
}