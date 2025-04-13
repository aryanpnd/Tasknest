package handlers

import (
	"github.com/aryanpnd/Tasknest/db"
	"github.com/aryanpnd/Tasknest/models"
	"github.com/gofiber/fiber/v2"
)

func RenderTodosPage(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	// Find the user
	var user models.User
	if err := db.DB.Preload("Todos").Where("username = ?", username).First(&user).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	return c.Render("todos", fiber.Map{
		"Username": user.Username,
		"Todos":    user.Todos,
	})
}
