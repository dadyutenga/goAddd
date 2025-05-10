package routes

import (
	"goAddd/internal/handlers"
	"goAddd/internal/middleware"
    "github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the API routes using Fiber
func SetupRoutes(app *fiber.App) {
	// Public routes
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// Protected routes (require authentication)
	api := app.Group("/api", middleware.AuthMiddleware())
	{
		// User management
		api.Get("/users", handlers.GetUsers)
		api.Get("/users/:id", handlers.GetUser)
		api.Put("/users/:id", handlers.UpdateUser)

		// Organization management
		api.Post("/organizations", handlers.CreateOrganization)
		api.Get("/organizations", handlers.GetOrganizations)

		// Kanban board
		api.Post("/boards", handlers.CreateBoard)
		api.Get("/boards", handlers.GetBoards)
		api.Post("/columns", handlers.CreateColumn)
		api.Post("/cards", handlers.CreateCard)
		api.Put("/cards/:id/move", handlers.MoveCard)

		// Chat
		api.Post("/messages", handlers.SendMessage)
		api.Get("/messages", handlers.GetMessages)
	}
}
