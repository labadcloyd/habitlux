package routes

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/controllers/habitControllers"
)

func HabitRoutes(app *fiber.App) {
	app.Get("/api/habit", controllers.GetAllUserHabits)
	app.Post("/api/habit", controllers.CreateHabit)
	app.Post("/api/habitlist", controllers.CreateHabitList)
	app.Put("/api/habit", controllers.UpdateHabit)
	app.Put("/api/habitlist", controllers.UpdateHabitList)
	app.Delete("/api/habit", controllers.DeleteHabit)
	app.Delete("/api/habitlist", controllers.DeleteHabitList)
}