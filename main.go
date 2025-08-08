package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	database "UBookTsk/Database"
	repository "UBookTsk/RepostoryPkg"
	"UBookTsk/api"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Initialize repository and handler
	bookRepo := repository.NewBookRepository(db)
	handler := api.NewHandler(bookRepo)

	// Setup Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(500).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger.New())

	// Register routes
	handler.RegisterRoutes(app)

	// Run server
	port := os.Getenv("PORT")
	log.Println("Server started on port", port)
	log.Fatal(app.Listen(":" + port))
}
