package main

import (
	"log"

	"github.com/aryanpnd/Tasknest/db"
	"github.com/aryanpnd/Tasknest/middlewares"
	"github.com/aryanpnd/Tasknest/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	db.Init()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.SetupUserRoutes(app)

	app.Use(middlewares.AuthRequired())

	routes.SetupTodoRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/todos/html", fiber.StatusSeeOther)
	})

	log.Fatal(app.Listen(":8080"))
}
