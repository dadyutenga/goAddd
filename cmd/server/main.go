package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"../internal/config"
	"../internal/db"
	"../internal/routes"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	if err := db.InitDB(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()

	// Set up Fiber app
	app := fiber.New()

	// Initialize routes
	routes.SetupRoutes(app)

	// Start server
	addr := ":" + cfg.ServerPort
	log.Printf("Server starting on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}