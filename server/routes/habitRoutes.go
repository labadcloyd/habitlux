package routes

import (
	controllers "habit-tracker/controllers/habitControllers"

	"github.com/gofiber/fiber/v2"
)

func HabitRoutes(app *fiber.App) {
	r := app.Group("/api")

	r.Get("/habit", controllers.GetAllUserHabits)
	r.Post("/habit", controllers.CreateHabit)
	r.Post("/habitlist", controllers.CreateHabitList)
	r.Put("/habit", controllers.UpdateHabit)
	r.Put("/habitlist", controllers.UpdateHabitList)
	r.Delete("/habit", controllers.DeleteHabit)
	r.Delete("/habitlist", controllers.DeleteHabitList)
}
