package handlers

import (
	"log"
	"strconv"

	"github.com/aryanpnd/Tasknest/db"
	"github.com/aryanpnd/Tasknest/models"
	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint) // jwt.MapClaims returns float64

	var todos []models.Todo
	if err := db.DB.Where("user_id = ?", uint(userID)).Find(&todos).Error; err != nil {
		log.Println("Error retrieving todos:", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var todo models.Todo
	if err := db.DB.Where("id = ? AND user_id = ?", id, uint(userID)).First(&todo).Error; err != nil {
		log.Println("Error retrieving todo:", err)
		return c.Status(404).SendString("Todo not found")
	}

	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	todo.UserID = uint(userID)

	if err := db.DB.Create(&todo).Error; err != nil {
		log.Println("Error creating todo:", err)
		return c.Status(500).SendString("Failed to create todo")
	}

	return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var todo models.Todo
	if err := db.DB.Where("id = ? AND user_id = ?", id, uint(userID)).First(&todo).Error; err != nil {
		log.Println("Error retrieving todo:", err)
		return c.Status(404).SendString("Todo not found")
	}

	// Apply updates
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	todo.UserID = uint(userID) // Ensure user cannot hijack another's todo

	if err := db.DB.Save(&todo).Error; err != nil {
		log.Println("Error updating todo:", err)
		return c.Status(500).SendString("Failed to update todo")
	}
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	todoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	// Ensure the todo belongs to the user
	if err := db.DB.Where("id = ? AND user_id = ?", todoID, uint(userID)).Delete(&models.Todo{}).Error; err != nil {
		log.Println("Error deleting todo:", err)
		return c.Status(500).SendString("Failed to delete todo")
	}

	return c.SendStatus(204)
}
