package routes

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/controllers/habitControllers"
)

func HabitRoutes(app *fiber.App) {
	app.Get("/api/habit", controllers.GetAllUserHabits)
	app.Post("/api/habit", controllers.CreateHabit)
	app.Put("/api/habit", controllers.UpdateHabit)
	app.Delete("/api/habit", controllers.DeleteHabit)
	app.Post("/api/habitList", controllers.CreateHabitList)
	app.Put("/api/habitList", controllers.UpdateHabitList)
	app.Delete("/api/habitList", controllers.DeleteHabitList)
}