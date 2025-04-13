// routes/todo_routes.go
package routes

import (
	"github.com/aryanpnd/Tasknest/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app *fiber.App) {
	app.Get("/todos/html", handlers.RenderTodosPage) // HTML view

	app.Get("/todos", handlers.GetTodos) // API
	app.Get("/todos/:id", handlers.GetTodo)
	app.Post("/todos", handlers.CreateTodo)
	app.Put("/todos/:id", handlers.UpdateTodo)
	app.Delete("/todos/:id", handlers.DeleteTodo)
}
