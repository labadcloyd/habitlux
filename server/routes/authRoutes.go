package routes

import (
	controllers "habit-tracker/controllers/authControllers"
	"habit-tracker/setup"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	r := app.Group("/api")

	r.Post("/signup", func(c *fiber.Ctx) error { return controllers.Signup(c, setup.DB) })
	r.Post("/login", func(c *fiber.Ctx) error { return controllers.Login(c, setup.DB) })
	r.Get("/user", func(c *fiber.Ctx) error { return controllers.User(c, setup.DB) })
	r.Post("/logout", func(c *fiber.Ctx) error { return controllers.Logout(c, setup.DB) })
	r.Get("/verifytoken", func(c *fiber.Ctx) error { return controllers.VerifyToken(c, setup.DB) })
	r.Post("/demologin", func(c *fiber.Ctx) error { return controllers.DemoLogin(c, setup.DB) })
}
