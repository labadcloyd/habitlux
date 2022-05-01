package routes

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/controllers/authControllers"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/signup", controllers.Signup)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/verifytoken", controllers.VerifyToken)
}