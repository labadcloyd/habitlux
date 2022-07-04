package main

import (
	"habit-tracker/helpers"
	"habit-tracker/routes"
	"habit-tracker/setup"
	"log"
)

func main() {
	setup.ConnectDB()
	app := setup.SetupApp()

	port := helpers.GoDotEnvVariable("PORT")
	if port == "" {
		port = "3000"
	}

	// ui
	routes.StaticRoutes(app)

	// routes
	routes.AuthRoutes(app)
	routes.HabitRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
