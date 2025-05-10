package handlers

import (
	"../db"
	"../models"
	"github.com/gofiber/fiber/v2"
)

// GetUsers retrieves a list of all users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if result := db.DB.Find(&users); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}
	return c.JSON(users)
}

// Register handles user registration
func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Hash password and save user to database (implementation omitted for brevity)
	// ...

	return c.Status(fiber.StatusCreated).JSON(user)
}
