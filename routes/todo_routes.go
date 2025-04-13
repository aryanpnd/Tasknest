package routes

import (
	"github.com/aryanpnd/Tasknest/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Get all todos
	app.Get("/todos", handlers.GetTodos)

	// Get a single todo by ID
	app.Get("/todos/:id", handlers.GetTodo)

	// Create a new todo
	app.Post("/todos", handlers.CreateTodo)

	// Update a todo
	app.Put("/todos/:id", handlers.UpdateTodo)

	// Delete a todo
	app.Delete("/todos/:id", handlers.DeleteTodo)
}
