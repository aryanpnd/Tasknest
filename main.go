package main

import (
	"log"

	"github.com/aryanpnd/Tasknest/db"
	"github.com/aryanpnd/Tasknest/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Init()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
