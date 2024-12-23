package main

import (
	"fmt"
	// "log"
	// "os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Use the logger middleware
	app.Use(logger.New())

	app.Use(cors.New())
	app.Use(helmet.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		healthcheck := map[string]interface{}{
			"uptime":    time.Now(),
			"message":   "OK",
			"timestamp": time.Now().Format(time.RFC3339),
		}

		return c.JSON(healthcheck)
	})
	
	fmt.Println("Server is running on port 3000", app.Listen(":3000"))

}

