package routes

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/controllers/habitControllers"
)

func HabitRoutes(app *fiber.App) {
	app.Post("/api/habit", controllers.CreateHabit)
	app.Get("/api/habit", controllers.GetAllHabits)
	// app.Post("/api/createHabitList", controllers.CreateHabitList)
	app.Get("/api/habitlists", controllers.GetAllHabitLists)
}