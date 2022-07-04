package routes

import (
	controllers "habit-tracker/controllers/habitControllers"
	"habit-tracker/setup"

	"github.com/gofiber/fiber/v2"
)

func HabitRoutes(app *fiber.App) {
	r := app.Group("/api")

	r.Get("/habit", func(c *fiber.Ctx) error { return controllers.GetAllUserHabits(c, setup.DB) })
	r.Post("/habit", func(c *fiber.Ctx) error { return controllers.CreateHabit(c, setup.DB) })
	r.Post("/habitlist", func(c *fiber.Ctx) error { return controllers.CreateHabitList(c, setup.DB) })
	r.Put("/habit", func(c *fiber.Ctx) error { return controllers.UpdateHabit(c, setup.DB) })
	r.Put("/habitlist", func(c *fiber.Ctx) error { return controllers.UpdateHabitList(c, setup.DB) })
	r.Delete("/habit", func(c *fiber.Ctx) error { return controllers.DeleteHabit(c, setup.DB) })
	r.Delete("/habitlist", func(c *fiber.Ctx) error { return controllers.DeleteHabitList(c, setup.DB) })
}
