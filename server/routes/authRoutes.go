package routes

import (
	controllers "habit-tracker/controllers/authControllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	r := app.Group("/api")

	r.Post("/signup", controllers.Signup)
	r.Post("/login", controllers.Login)
	r.Get("/user", controllers.User)
	r.Post("/logout", controllers.Logout)
	r.Get("/verifytoken", controllers.VerifyToken)
	r.Post("/demologin", controllers.DemoLogin)
}
