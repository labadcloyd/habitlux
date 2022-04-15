package routes

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/controllers/habitControllers"
)

func HabitRoutes(app *fiber.App) {
	app.Post("/api/habit", controllers.CreateHabit)
	app.Get("/api/habit", controllers.GetAllHabits)
}