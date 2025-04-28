package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/makalin/Sayo/handlers"
	"github.com/makalin/Sayo/models"
)

func main() {
	// Initialize database
	db, err := models.NewDatabase("database/sayo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize handler
	handler := handlers.NewHandler(db)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Sayo",
	})

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// Public routes
	app.Post("/api/register", handler.Register)
	app.Post("/api/login", handler.Login)

	// Protected routes
	api := app.Group("/api", handler.AuthMiddleware)
	api.Post("/says", handler.CreateSay)
	api.Get("/says", handler.GetSays)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
