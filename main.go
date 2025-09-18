package main

import (
	"log"

	_ "backend/docs" // Import generated docs with correct module path

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Backend API
// @version 1.0
// @description This is a sample backend API with authentication
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:3000
// @BasePath /
func main() {
	// Initialize database
	InitDatabase()

	// Create fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello world",
		})
	})

	// Authentication routes
	auth := app.Group("/auth")
	auth.Post("/register", Register)
	auth.Post("/login", Login)

	// Swagger route - Fixed configuration
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Start server on port 3000
	log.Fatal(app.Listen(":3000"))
}
