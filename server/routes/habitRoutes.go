package routes

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/controllers/habitControllers"
)

func HabitRoutes(app *fiber.App) {
	app.Get("/api/habit", controllers.GetAllUserHabits)
	app.Post("/api/habit", controllers.CreateHabit)
	app.Post("/api/habitList", controllers.CreateHabitList)
}