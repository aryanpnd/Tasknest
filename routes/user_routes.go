package routes

import (
	"github.com/aryanpnd/Tasknest/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	// app.Get("/users", handlers.GetUsers)

	app.Get("/register", handlers.ShowRegisterPage)
	app.Post("/register", handlers.RegisterUser)

	app.Get("/login", handlers.ShowLoginPage)
	app.Post("/login", handlers.LoginUser)
}
