package handlers

import (
	"log"

	"github.com/aryanpnd/Tasknest/db"
	"github.com/aryanpnd/Tasknest/models"
	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	if err := db.DB.Find(&todos).Error; err != nil {
		log.Println("Error retrieving todos:", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		log.Println("Error retrieving todo:", err)
		return c.Status(404).SendString("Todo not found")
	}

	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	if err := db.DB.Create(&todo).Error; err != nil {
		log.Println("Error creating todo:", err)
		return c.Status(500).SendString("Failed to create todo")
	}

	return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		log.Println("Error retrieving todo:", err)
		return c.Status(404).SendString("Todo not found")
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	if err := db.DB.Save(&todo).Error; err != nil {
		log.Println("Error updating todo:", err)
		return c.Status(500).SendString("Failed to update todo")
	}
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := db.DB.Delete(&models.Todo{}, id).Error; err != nil {
		log.Println("Error deleting todo:", err)
		return c.Status(500).SendString("Failed to delete todo")
	}
	return c.SendStatus(204)
}
